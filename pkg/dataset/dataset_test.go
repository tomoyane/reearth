package dataset

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestDataset_Interface(t *testing.T) {
	f1 := id.NewDatasetSchemaFieldID()
	f2 := id.NewDatasetSchemaFieldID()
	sid := id.NewDatasetSchemaID()

	tests := []struct {
		name    string
		schema  *Schema
		dataset *Dataset
		want    map[string]interface{}
	}{
		{
			name: "ok",
			schema: NewSchema().ID(sid).Scene(id.NewSceneID()).Fields([]*SchemaField{
				NewSchemaField().ID(f1).Name("foo").Type(ValueTypeNumber).MustBuild(),
				NewSchemaField().ID(f2).Name("bar").Type(ValueTypeLatLng).MustBuild(),
			}).MustBuild(),
			dataset: New().NewID().Scene(id.NewSceneID()).Schema(sid).Fields([]*Field{
				NewField(f1, ValueTypeNumber.ValueFrom(1), ""),
				NewField(f2, ValueTypeLatLng.ValueFrom(LatLng{Lat: 1, Lng: 2}), ""),
			}).MustBuild(),
			want: map[string]interface{}{
				"foo": float64(1),
				"bar": LatLng{Lat: 1, Lng: 2},
			},
		},
		{
			name:    "empty",
			dataset: &Dataset{},
		},
		{
			name: "nil",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dataset.Interface(tt.schema))
		})
	}
}

func TestDataset_InterfaceWithFieldIDs(t *testing.T) {
	f1 := id.NewDatasetSchemaFieldID()
	f2 := id.NewDatasetSchemaFieldID()

	tests := []struct {
		name    string
		dataset *Dataset
		want    map[string]interface{}
	}{
		{
			name: "ok",
			dataset: New().NewID().Scene(id.NewSceneID()).Schema(id.NewDatasetSchemaID()).Fields([]*Field{
				NewField(f1, ValueTypeNumber.ValueFrom(1), ""),
				NewField(f2, ValueTypeLatLng.ValueFrom(LatLng{Lat: 1, Lng: 2}), ""),
			}).MustBuild(),
			want: map[string]interface{}{
				f1.String(): float64(1),
				f2.String(): LatLng{Lat: 1, Lng: 2},
			},
		},
		{
			name:    "empty",
			dataset: &Dataset{},
			want:    map[string]interface{}{},
		},
		{
			name: "nil",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.dataset.InterfaceWithFieldIDs())
		})
	}
}
