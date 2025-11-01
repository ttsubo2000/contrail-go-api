//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	config_properties_properties = iota
	config_properties_id_perms
	config_properties_perms2
	config_properties_annotations
	config_properties_display_name
	config_properties_tag_refs
	config_properties_max_
)

type ConfigProperties struct {
        contrail.ObjectBase
	properties KeyValuePairs
	id_perms IdPermsType
	perms2 PermType2
	annotations KeyValuePairs
	display_name string
	tag_refs contrail.ReferenceList
        valid [config_properties_max_] bool
        modified [config_properties_max_] bool
        baseMap map[string]contrail.ReferenceList
}

func (obj *ConfigProperties) GetType() string {
        return "config-properties"
}

func (obj *ConfigProperties) GetDefaultParent() []string {
        name := []string{"default-global-system-config"}
        return name
}

func (obj *ConfigProperties) GetDefaultParentType() string {
        return "global-system-config"
}

func (obj *ConfigProperties) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *ConfigProperties) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *ConfigProperties) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *ConfigProperties) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *ConfigProperties) UpdateDone() {
        for i := range obj.modified { obj.modified[i] = false }
        obj.baseMap = nil
}


func (obj *ConfigProperties) GetProperties() KeyValuePairs {
        return obj.properties
}

func (obj *ConfigProperties) SetProperties(value *KeyValuePairs) {
        obj.properties = *value
        obj.modified[config_properties_properties] = true
}

func (obj *ConfigProperties) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *ConfigProperties) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified[config_properties_id_perms] = true
}

func (obj *ConfigProperties) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *ConfigProperties) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified[config_properties_perms2] = true
}

func (obj *ConfigProperties) GetAnnotations() KeyValuePairs {
        return obj.annotations
}

func (obj *ConfigProperties) SetAnnotations(value *KeyValuePairs) {
        obj.annotations = *value
        obj.modified[config_properties_annotations] = true
}

func (obj *ConfigProperties) GetDisplayName() string {
        return obj.display_name
}

func (obj *ConfigProperties) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified[config_properties_display_name] = true
}

func (obj *ConfigProperties) readTagRefs() error {
        if !obj.IsTransient() &&
                !obj.valid[config_properties_tag_refs] {
                err := obj.GetField(obj, "tag_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ConfigProperties) GetTagRefs() (
        contrail.ReferenceList, error) {
        err := obj.readTagRefs()
        if err != nil {
                return nil, err
        }
        return obj.tag_refs, nil
}

func (obj *ConfigProperties) AddTag(
        rhs *Tag) error {
        err := obj.readTagRefs()
        if err != nil {
                return err
        }

        if !obj.modified[config_properties_tag_refs] {
                obj.storeReferenceBase("tag", obj.tag_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.tag_refs = append(obj.tag_refs, ref)
        obj.modified[config_properties_tag_refs] = true
        return nil
}

func (obj *ConfigProperties) DeleteTag(uuid string) error {
        err := obj.readTagRefs()
        if err != nil {
                return err
        }

        if !obj.modified[config_properties_tag_refs] {
                obj.storeReferenceBase("tag", obj.tag_refs)
        }

        for i, ref := range obj.tag_refs {
                if ref.Uuid == uuid {
                        obj.tag_refs = append(
                                obj.tag_refs[:i],
                                obj.tag_refs[i+1:]...)
                        break
                }
        }
        obj.modified[config_properties_tag_refs] = true
        return nil
}

func (obj *ConfigProperties) ClearTag() {
        if obj.valid[config_properties_tag_refs] &&
           !obj.modified[config_properties_tag_refs] {
                obj.storeReferenceBase("tag", obj.tag_refs)
        }
        obj.tag_refs = make([]contrail.Reference, 0)
        obj.valid[config_properties_tag_refs] = true
        obj.modified[config_properties_tag_refs] = true
}

func (obj *ConfigProperties) SetTagList(
        refList []contrail.ReferencePair) {
        obj.ClearTag()
        obj.tag_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.tag_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *ConfigProperties) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified[config_properties_properties] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.properties)
                if err != nil {
                        return nil, err
                }
                msg["properties"] = &value
        }

        if obj.modified[config_properties_id_perms] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified[config_properties_perms2] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified[config_properties_annotations] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.annotations)
                if err != nil {
                        return nil, err
                }
                msg["annotations"] = &value
        }

        if obj.modified[config_properties_display_name] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if len(obj.tag_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.tag_refs)
                if err != nil {
                        return nil, err
                }
                msg["tag_refs"] = &value
        }

        return json.Marshal(msg)
}

func (obj *ConfigProperties) UnmarshalJSON(body []byte) error {
        var m map[string]json.RawMessage
        err := json.Unmarshal(body, &m)
        if err != nil {
                return err
        }
        err = obj.UnmarshalCommon(m)
        if err != nil {
                return err
        }
        for key, value := range m {
                switch key {
                case "properties":
                        err = json.Unmarshal(value, &obj.properties)
                        if err == nil {
                                obj.valid[config_properties_properties] = true
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid[config_properties_id_perms] = true
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid[config_properties_perms2] = true
                        }
                        break
                case "annotations":
                        err = json.Unmarshal(value, &obj.annotations)
                        if err == nil {
                                obj.valid[config_properties_annotations] = true
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid[config_properties_display_name] = true
                        }
                        break
                case "tag_refs":
                        err = json.Unmarshal(value, &obj.tag_refs)
                        if err == nil {
                                obj.valid[config_properties_tag_refs] = true
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ConfigProperties) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified[config_properties_properties] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.properties)
                if err != nil {
                        return nil, err
                }
                msg["properties"] = &value
        }

        if obj.modified[config_properties_id_perms] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified[config_properties_perms2] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified[config_properties_annotations] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.annotations)
                if err != nil {
                        return nil, err
                }
                msg["annotations"] = &value
        }

        if obj.modified[config_properties_display_name] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified[config_properties_tag_refs] {
                if len(obj.tag_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["tag_refs"] = &value
                } else if !obj.hasReferenceBase("tag") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.tag_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["tag_refs"] = &value
                }
        }


        return json.Marshal(msg)
}

func (obj *ConfigProperties) UpdateReferences() error {

        if obj.modified[config_properties_tag_refs] &&
           len(obj.tag_refs) > 0 &&
           obj.hasReferenceBase("tag") {
                err := obj.UpdateReference(
                        obj, "tag",
                        obj.tag_refs,
                        obj.baseMap["tag"])
                if err != nil {
                        return err
                }
        }

        return nil
}

func ConfigPropertiesByName(c contrail.ApiClient, fqn string) (*ConfigProperties, error) {
    obj, err := c.FindByName("config-properties", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*ConfigProperties), nil
}

func ConfigPropertiesByUuid(c contrail.ApiClient, uuid string) (*ConfigProperties, error) {
    obj, err := c.FindByUuid("config-properties", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*ConfigProperties), nil
}
