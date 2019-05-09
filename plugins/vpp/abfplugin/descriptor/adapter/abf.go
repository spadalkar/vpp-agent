// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	"github.com/gogo/protobuf/proto"
	. "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	"github.com/ligato/vpp-agent/plugins/vpp/abfplugin/abfidx"
	"github.com/ligato/vpp-agent/api/models/vpp/abf"
)

////////// type-safe key-value pair with metadata //////////

type ABFKVWithMetadata struct {
	Key      string
	Value    *vpp_abf.ABF
	Metadata *abfidx.ABFMetadata
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type ABFDescriptor struct {
	Name                 string
	KeySelector          KeySelector
	ValueTypeName        string
	KeyLabel             func(key string) string
	ValueComparator      func(key string, oldValue, newValue *vpp_abf.ABF) bool
	NBKeyPrefix          string
	WithMetadata         bool
	MetadataMapFactory   MetadataMapFactory
	Validate             func(key string, value *vpp_abf.ABF) error
	Create               func(key string, value *vpp_abf.ABF) (metadata *abfidx.ABFMetadata, err error)
	Delete               func(key string, value *vpp_abf.ABF, metadata *abfidx.ABFMetadata) error
	Update               func(key string, oldValue, newValue *vpp_abf.ABF, oldMetadata *abfidx.ABFMetadata) (newMetadata *abfidx.ABFMetadata, err error)
	UpdateWithRecreate   func(key string, oldValue, newValue *vpp_abf.ABF, metadata *abfidx.ABFMetadata) bool
	Retrieve             func(correlate []ABFKVWithMetadata) ([]ABFKVWithMetadata, error)
	IsRetriableFailure   func(err error) bool
	DerivedValues        func(key string, value *vpp_abf.ABF) []KeyValuePair
	Dependencies         func(key string, value *vpp_abf.ABF) []Dependency
	RetrieveDependencies []string /* descriptor name */
}

////////// Descriptor adapter //////////

type ABFDescriptorAdapter struct {
	descriptor *ABFDescriptor
}

func NewABFDescriptor(typedDescriptor *ABFDescriptor) *KVDescriptor {
	adapter := &ABFDescriptorAdapter{descriptor: typedDescriptor}
	descriptor := &KVDescriptor{
		Name:                 typedDescriptor.Name,
		KeySelector:          typedDescriptor.KeySelector,
		ValueTypeName:        typedDescriptor.ValueTypeName,
		KeyLabel:             typedDescriptor.KeyLabel,
		NBKeyPrefix:          typedDescriptor.NBKeyPrefix,
		WithMetadata:         typedDescriptor.WithMetadata,
		MetadataMapFactory:   typedDescriptor.MetadataMapFactory,
		IsRetriableFailure:   typedDescriptor.IsRetriableFailure,
		RetrieveDependencies: typedDescriptor.RetrieveDependencies,
	}
	if typedDescriptor.ValueComparator != nil {
		descriptor.ValueComparator = adapter.ValueComparator
	}
	if typedDescriptor.Validate != nil {
		descriptor.Validate = adapter.Validate
	}
	if typedDescriptor.Create != nil {
		descriptor.Create = adapter.Create
	}
	if typedDescriptor.Delete != nil {
		descriptor.Delete = adapter.Delete
	}
	if typedDescriptor.Update != nil {
		descriptor.Update = adapter.Update
	}
	if typedDescriptor.UpdateWithRecreate != nil {
		descriptor.UpdateWithRecreate = adapter.UpdateWithRecreate
	}
	if typedDescriptor.Retrieve != nil {
		descriptor.Retrieve = adapter.Retrieve
	}
	if typedDescriptor.Dependencies != nil {
		descriptor.Dependencies = adapter.Dependencies
	}
	if typedDescriptor.DerivedValues != nil {
		descriptor.DerivedValues = adapter.DerivedValues
	}
	return descriptor
}

func (da *ABFDescriptorAdapter) ValueComparator(key string, oldValue, newValue proto.Message) bool {
	typedOldValue, err1 := castABFValue(key, oldValue)
	typedNewValue, err2 := castABFValue(key, newValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedOldValue, typedNewValue)
}

func (da *ABFDescriptorAdapter) Validate(key string, value proto.Message) (err error) {
	typedValue, err := castABFValue(key, value)
	if err != nil {
		return err
	}
	return da.descriptor.Validate(key, typedValue)
}

func (da *ABFDescriptorAdapter) Create(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castABFValue(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Create(key, typedValue)
}

func (da *ABFDescriptorAdapter) Update(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castABFValue(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castABFValue(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castABFMetadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Update(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *ABFDescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castABFValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castABFMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *ABFDescriptorAdapter) UpdateWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castABFValue(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castABFValue(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castABFMetadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.UpdateWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *ABFDescriptorAdapter) Retrieve(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []ABFKVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castABFValue(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castABFMetadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			ABFKVWithMetadata{
				Key:      kvpair.Key,
				Value:    typedValue,
				Metadata: typedMetadata,
				Origin:   kvpair.Origin,
			})
	}

	typedValues, err := da.descriptor.Retrieve(correlateWithType)
	if err != nil {
		return nil, err
	}
	var values []KVWithMetadata
	for _, typedKVWithMetadata := range typedValues {
		kvWithMetadata := KVWithMetadata{
			Key:      typedKVWithMetadata.Key,
			Metadata: typedKVWithMetadata.Metadata,
			Origin:   typedKVWithMetadata.Origin,
		}
		kvWithMetadata.Value = typedKVWithMetadata.Value
		values = append(values, kvWithMetadata)
	}
	return values, err
}

func (da *ABFDescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castABFValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *ABFDescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castABFValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

////////// Helper methods //////////

func castABFValue(key string, value proto.Message) (*vpp_abf.ABF, error) {
	typedValue, ok := value.(*vpp_abf.ABF)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castABFMetadata(key string, metadata Metadata) (*abfidx.ABFMetadata, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(*abfidx.ABFMetadata)
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}
