package storage

import (
	"encoding/json"
	"fmt"
	"gomarket/internal/errs"
	"os"

	"golang.org/x/exp/slices"
)

type Raw json.RawMessage

type Collection interface {
	// GetNextCode will retrieve the code you need to use on a new registry
	GetNextCode() int
	// Get will try to retrieve a single registry from the storage
	Get(code int, dest interface{}) error
	// List will try to retrieve all registries from the storage. You'll have to decode the element later using the DecodeRaw method
	List(offset int, limit int) ([]Raw, error)
	// DecodeRaw will decode an element data retrieved by the List method
	DecodeRaw(raw Raw, dest interface{}) error
	// Save will try to insert or update a single registry to the storage, overriding it if it already exists
	Save(code int, data interface{}) error
	// Delete will try to delete a single registry from the storage
	Delete(code int) error
}

type collection struct {
	Name     string      `json:"name"`
	Codes    []int       `json:"codes"`
	NextCode int         `json:"next_code"`
	storage  JsonStorage `json:"-"`
}

func NewCollection(storage JsonStorage, name string) (Collection, error) {
	c := &collection{
		Name:     name,
		Codes:    make([]int, 0),
		NextCode: 1,
		storage:  storage,
	}

	err := c.loadFromStorage()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *collection) GetNextCode() int {
	return c.NextCode
}

func (c *collection) Get(code int, dest interface{}) error {
	if !slices.Contains(c.Codes, code) {
		return errs.RegistryNotFoundErr
	}

	err := c.storage.Read(c.codePath(code), dest)
	return err
}

func (c *collection) List(offset int, limit int) ([]Raw, error) {
	output := make([]Raw, 0)
	for index := offset; index < offset+limit && index < len(c.Codes); index++ {
		var model Raw
		err := c.storage.Read(
			c.codePath(c.Codes[index]),
			&model,
		)
		if err != nil {
			return nil, err
		}
		output = append(output, model)
	}

	return output, nil
}

func (c *collection) DecodeRaw(raw Raw, dest interface{}) error {
	return json.Unmarshal(raw, dest)
}

func (c *collection) Save(code int, data interface{}) error {
	err := c.storage.Write(
		c.codePath(code),
		data,
	)
	if err != nil {
		return err
	}

	if !slices.Contains(c.Codes, code) {
		c.Codes = append(c.Codes, code)
		if code >= c.NextCode {
			c.NextCode = code + 1
		}
		c.sync()
	}

	return nil
}

func (c *collection) Delete(code int) error {
	for index, iCode := range c.Codes {
		if iCode != code {
			continue
		}

		err := c.storage.Delete(c.codePath(code))
		if err != nil {
			return err
		}

		c.Codes = slices.Delete(c.Codes, index, index+1)
		c.sync()
		return nil
	}

	return errs.RegistryNotFoundErr
}

// codePath retrieves the filename of a registry, following the pattern "collection_{name}_registry_{code}.json"
func (c *collection) codePath(code int) string {
	return fmt.Sprintf("collection_%s_registry_%d.json", c.Name, code)
}

// collectionPath retrieves the filename of a collection, following the pattern "collection_{name}.json"
func (c *collection) collectionPath() string {
	return fmt.Sprintf("collection_%s.json", c.Name)
}

// loadFromStorage will try to retrieve the stored collection data from storage
func (c *collection) loadFromStorage() error {
	var data collection
	err := c.storage.Read(
		c.collectionPath(),
		&data,
	)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	c.Codes = data.Codes
	c.NextCode = data.NextCode
	return nil
}

// sync will try to store the collection data on the storage
func (c *collection) sync() {
	err := c.storage.Write(
		c.collectionPath(),
		c,
	)
	if err != nil {
		panic(err)
	}
}
