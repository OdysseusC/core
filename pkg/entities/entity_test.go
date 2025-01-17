package entities

import (
	"context"
	"fmt"
	"testing"

	ants "github.com/panjf2000/ants/v2"
	"github.com/stretchr/testify/assert"
)

func TestEntity(t *testing.T) {
	coroutinePool, err := ants.NewPool(500)
	if nil != err {
		panic(err)
	}

	mgr, _ := NewEntityManager(context.Background(), coroutinePool)

	enty1, err1 := newEntity(context.Background(), mgr, &EntityBase{ID: "", PluginID: "abcd", Owner: "tomas", Version: 001})
	enty2, err2 := newEntity(context.Background(), mgr, &EntityBase{ID: "", PluginID: "abcd", Owner: "tomas", Version: 001})

	t.Log(enty1, enty2, err1, err2)
}

func TestGetProperties(t *testing.T) {
	coroutinePool, err := ants.NewPool(500)
	if nil != err {
		panic(err)
	}

	mgr, _ := NewEntityManager(context.Background(), coroutinePool)

	entity, err := newEntity(context.Background(), mgr, &EntityBase{ID: "", PluginID: "abcd", Owner: "tomas", Version: 001})
	if nil != err {
		panic(err)
	}

	entity.SetProperties(&EntityBase{
		KValues: map[string]interface{}{
			"temp1":   15,
			"temp2":   23,
			"light":   555,
			"say":     "hello",
			"friends": []string{"tom", "tony"},
			"user": map[string]interface{}{
				"name": "john",
				"age":  20,
			},
		},
	})

	eb := entity.GetAllProperties()
	t.Log(eb.KValues)

	// delete some field.
	delete(eb.KValues, "light")

	props1 := entity.GetAllProperties()
	t.Log(props1)
}

func TestEntity_getEntity(t *testing.T) {
	coroutinePool, err := ants.NewPool(500)
	assert.Nil(t, err)

	mgr, _ := NewEntityManager(context.Background(), coroutinePool)

	en, err := newEntity(context.Background(), mgr, &EntityBase{ID: "", PluginID: "abcd", Owner: "tomas", Version: 001})
	assert.Nil(t, err)
	ce := en.getEntityBase()

	assert.NotEqual(t, fmt.Sprintf("%p", en), fmt.Sprintf("%p", &ce))
}
