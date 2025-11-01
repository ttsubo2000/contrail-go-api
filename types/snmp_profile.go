//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/ttsubo2000/contrail-go-api"
)

const (
	snmp_profile_snmp_profile_is_default = iota
	snmp_profile_snmp_parameters
	snmp_profile_id_perms
	snmp_profile_perms2
	snmp_profile_annotations
	snmp_profile_display_name
	snmp_profile_tag_refs
	snmp_profile_telemetry_profile_back_refs
	snmp_profile_max_
)

type SnmpProfile struct {
        contrail.ObjectBase
	snmp_profile_is_default bool
	snmp_parameters SnmpParameters
	id_perms IdPermsType
	perms2 PermType2
	annotations KeyValuePairs
	display_name string
	tag_refs contrail.ReferenceList
	telemetry_profile_back_refs contrail.ReferenceList
        valid [snmp_profile_max_] bool
        modified [snmp_profile_max_] bool
        baseMap map[string]contrail.ReferenceList
}

func (obj *SnmpProfile) GetType() string {
        return "snmp-profile"
}

func (obj *SnmpProfile) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project"}
        return name
}

func (obj *SnmpProfile) GetDefaultParentType() string {
        return "project"
}

func (obj *SnmpProfile) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *SnmpProfile) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *SnmpProfile) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *SnmpProfile) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *SnmpProfile) UpdateDone() {
        for i := range obj.modified { obj.modified[i] = false }
        obj.baseMap = nil
}


func (obj *SnmpProfile) GetSnmpProfileIsDefault() bool {
        return obj.snmp_profile_is_default
}

func (obj *SnmpProfile) SetSnmpProfileIsDefault(value bool) {
        obj.snmp_profile_is_default = value
        obj.modified[snmp_profile_snmp_profile_is_default] = true
}

func (obj *SnmpProfile) GetSnmpParameters() SnmpParameters {
        return obj.snmp_parameters
}

func (obj *SnmpProfile) SetSnmpParameters(value *SnmpParameters) {
        obj.snmp_parameters = *value
        obj.modified[snmp_profile_snmp_parameters] = true
}

func (obj *SnmpProfile) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *SnmpProfile) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified[snmp_profile_id_perms] = true
}

func (obj *SnmpProfile) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *SnmpProfile) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified[snmp_profile_perms2] = true
}

func (obj *SnmpProfile) GetAnnotations() KeyValuePairs {
        return obj.annotations
}

func (obj *SnmpProfile) SetAnnotations(value *KeyValuePairs) {
        obj.annotations = *value
        obj.modified[snmp_profile_annotations] = true
}

func (obj *SnmpProfile) GetDisplayName() string {
        return obj.display_name
}

func (obj *SnmpProfile) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified[snmp_profile_display_name] = true
}

func (obj *SnmpProfile) readTagRefs() error {
        if !obj.IsTransient() &&
                !obj.valid[snmp_profile_tag_refs] {
                err := obj.GetField(obj, "tag_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *SnmpProfile) GetTagRefs() (
        contrail.ReferenceList, error) {
        err := obj.readTagRefs()
        if err != nil {
                return nil, err
        }
        return obj.tag_refs, nil
}

func (obj *SnmpProfile) AddTag(
        rhs *Tag) error {
        err := obj.readTagRefs()
        if err != nil {
                return err
        }

        if !obj.modified[snmp_profile_tag_refs] {
                obj.storeReferenceBase("tag", obj.tag_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.tag_refs = append(obj.tag_refs, ref)
        obj.modified[snmp_profile_tag_refs] = true
        return nil
}

func (obj *SnmpProfile) DeleteTag(uuid string) error {
        err := obj.readTagRefs()
        if err != nil {
                return err
        }

        if !obj.modified[snmp_profile_tag_refs] {
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
        obj.modified[snmp_profile_tag_refs] = true
        return nil
}

func (obj *SnmpProfile) ClearTag() {
        if obj.valid[snmp_profile_tag_refs] &&
           !obj.modified[snmp_profile_tag_refs] {
                obj.storeReferenceBase("tag", obj.tag_refs)
        }
        obj.tag_refs = make([]contrail.Reference, 0)
        obj.valid[snmp_profile_tag_refs] = true
        obj.modified[snmp_profile_tag_refs] = true
}

func (obj *SnmpProfile) SetTagList(
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


func (obj *SnmpProfile) readTelemetryProfileBackRefs() error {
        if !obj.IsTransient() &&
                !obj.valid[snmp_profile_telemetry_profile_back_refs] {
                err := obj.GetField(obj, "telemetry_profile_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *SnmpProfile) GetTelemetryProfileBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readTelemetryProfileBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.telemetry_profile_back_refs, nil
}

func (obj *SnmpProfile) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified[snmp_profile_snmp_profile_is_default] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.snmp_profile_is_default)
                if err != nil {
                        return nil, err
                }
                msg["snmp_profile_is_default"] = &value
        }

        if obj.modified[snmp_profile_snmp_parameters] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.snmp_parameters)
                if err != nil {
                        return nil, err
                }
                msg["snmp_parameters"] = &value
        }

        if obj.modified[snmp_profile_id_perms] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified[snmp_profile_perms2] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified[snmp_profile_annotations] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.annotations)
                if err != nil {
                        return nil, err
                }
                msg["annotations"] = &value
        }

        if obj.modified[snmp_profile_display_name] {
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

func (obj *SnmpProfile) UnmarshalJSON(body []byte) error {
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
                case "snmp_profile_is_default":
                        err = json.Unmarshal(value, &obj.snmp_profile_is_default)
                        if err == nil {
                                obj.valid[snmp_profile_snmp_profile_is_default] = true
                        }
                        break
                case "snmp_parameters":
                        err = json.Unmarshal(value, &obj.snmp_parameters)
                        if err == nil {
                                obj.valid[snmp_profile_snmp_parameters] = true
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid[snmp_profile_id_perms] = true
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid[snmp_profile_perms2] = true
                        }
                        break
                case "annotations":
                        err = json.Unmarshal(value, &obj.annotations)
                        if err == nil {
                                obj.valid[snmp_profile_annotations] = true
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid[snmp_profile_display_name] = true
                        }
                        break
                case "tag_refs":
                        err = json.Unmarshal(value, &obj.tag_refs)
                        if err == nil {
                                obj.valid[snmp_profile_tag_refs] = true
                        }
                        break
                case "telemetry_profile_back_refs":
                        err = json.Unmarshal(value, &obj.telemetry_profile_back_refs)
                        if err == nil {
                                obj.valid[snmp_profile_telemetry_profile_back_refs] = true
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *SnmpProfile) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified[snmp_profile_snmp_profile_is_default] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.snmp_profile_is_default)
                if err != nil {
                        return nil, err
                }
                msg["snmp_profile_is_default"] = &value
        }

        if obj.modified[snmp_profile_snmp_parameters] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.snmp_parameters)
                if err != nil {
                        return nil, err
                }
                msg["snmp_parameters"] = &value
        }

        if obj.modified[snmp_profile_id_perms] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified[snmp_profile_perms2] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified[snmp_profile_annotations] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.annotations)
                if err != nil {
                        return nil, err
                }
                msg["annotations"] = &value
        }

        if obj.modified[snmp_profile_display_name] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified[snmp_profile_tag_refs] {
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

func (obj *SnmpProfile) UpdateReferences() error {

        if obj.modified[snmp_profile_tag_refs] &&
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

func SnmpProfileByName(c contrail.ApiClient, fqn string) (*SnmpProfile, error) {
    obj, err := c.FindByName("snmp-profile", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*SnmpProfile), nil
}

func SnmpProfileByUuid(c contrail.ApiClient, uuid string) (*SnmpProfile, error) {
    obj, err := c.FindByUuid("snmp-profile", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*SnmpProfile), nil
}
