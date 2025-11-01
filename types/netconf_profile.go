//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/ttsubo2000/contrail-go-api"
)

const (
	netconf_profile_netconf_profile_is_default = iota
	netconf_profile_netconf_parameters
	netconf_profile_id_perms
	netconf_profile_perms2
	netconf_profile_annotations
	netconf_profile_display_name
	netconf_profile_tag_refs
	netconf_profile_telemetry_profile_back_refs
	netconf_profile_max_
)

type NetconfProfile struct {
        contrail.ObjectBase
	netconf_profile_is_default bool
	netconf_parameters NetconfParameters
	id_perms IdPermsType
	perms2 PermType2
	annotations KeyValuePairs
	display_name string
	tag_refs contrail.ReferenceList
	telemetry_profile_back_refs contrail.ReferenceList
        valid [netconf_profile_max_] bool
        modified [netconf_profile_max_] bool
        baseMap map[string]contrail.ReferenceList
}

func (obj *NetconfProfile) GetType() string {
        return "netconf-profile"
}

func (obj *NetconfProfile) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project"}
        return name
}

func (obj *NetconfProfile) GetDefaultParentType() string {
        return "project"
}

func (obj *NetconfProfile) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *NetconfProfile) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *NetconfProfile) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *NetconfProfile) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *NetconfProfile) UpdateDone() {
        for i := range obj.modified { obj.modified[i] = false }
        obj.baseMap = nil
}


func (obj *NetconfProfile) GetNetconfProfileIsDefault() bool {
        return obj.netconf_profile_is_default
}

func (obj *NetconfProfile) SetNetconfProfileIsDefault(value bool) {
        obj.netconf_profile_is_default = value
        obj.modified[netconf_profile_netconf_profile_is_default] = true
}

func (obj *NetconfProfile) GetNetconfParameters() NetconfParameters {
        return obj.netconf_parameters
}

func (obj *NetconfProfile) SetNetconfParameters(value *NetconfParameters) {
        obj.netconf_parameters = *value
        obj.modified[netconf_profile_netconf_parameters] = true
}

func (obj *NetconfProfile) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *NetconfProfile) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified[netconf_profile_id_perms] = true
}

func (obj *NetconfProfile) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *NetconfProfile) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified[netconf_profile_perms2] = true
}

func (obj *NetconfProfile) GetAnnotations() KeyValuePairs {
        return obj.annotations
}

func (obj *NetconfProfile) SetAnnotations(value *KeyValuePairs) {
        obj.annotations = *value
        obj.modified[netconf_profile_annotations] = true
}

func (obj *NetconfProfile) GetDisplayName() string {
        return obj.display_name
}

func (obj *NetconfProfile) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified[netconf_profile_display_name] = true
}

func (obj *NetconfProfile) readTagRefs() error {
        if !obj.IsTransient() &&
                !obj.valid[netconf_profile_tag_refs] {
                err := obj.GetField(obj, "tag_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *NetconfProfile) GetTagRefs() (
        contrail.ReferenceList, error) {
        err := obj.readTagRefs()
        if err != nil {
                return nil, err
        }
        return obj.tag_refs, nil
}

func (obj *NetconfProfile) AddTag(
        rhs *Tag) error {
        err := obj.readTagRefs()
        if err != nil {
                return err
        }

        if !obj.modified[netconf_profile_tag_refs] {
                obj.storeReferenceBase("tag", obj.tag_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.tag_refs = append(obj.tag_refs, ref)
        obj.modified[netconf_profile_tag_refs] = true
        return nil
}

func (obj *NetconfProfile) DeleteTag(uuid string) error {
        err := obj.readTagRefs()
        if err != nil {
                return err
        }

        if !obj.modified[netconf_profile_tag_refs] {
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
        obj.modified[netconf_profile_tag_refs] = true
        return nil
}

func (obj *NetconfProfile) ClearTag() {
        if obj.valid[netconf_profile_tag_refs] &&
           !obj.modified[netconf_profile_tag_refs] {
                obj.storeReferenceBase("tag", obj.tag_refs)
        }
        obj.tag_refs = make([]contrail.Reference, 0)
        obj.valid[netconf_profile_tag_refs] = true
        obj.modified[netconf_profile_tag_refs] = true
}

func (obj *NetconfProfile) SetTagList(
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


func (obj *NetconfProfile) readTelemetryProfileBackRefs() error {
        if !obj.IsTransient() &&
                !obj.valid[netconf_profile_telemetry_profile_back_refs] {
                err := obj.GetField(obj, "telemetry_profile_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *NetconfProfile) GetTelemetryProfileBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readTelemetryProfileBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.telemetry_profile_back_refs, nil
}

func (obj *NetconfProfile) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified[netconf_profile_netconf_profile_is_default] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.netconf_profile_is_default)
                if err != nil {
                        return nil, err
                }
                msg["netconf_profile_is_default"] = &value
        }

        if obj.modified[netconf_profile_netconf_parameters] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.netconf_parameters)
                if err != nil {
                        return nil, err
                }
                msg["netconf_parameters"] = &value
        }

        if obj.modified[netconf_profile_id_perms] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified[netconf_profile_perms2] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified[netconf_profile_annotations] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.annotations)
                if err != nil {
                        return nil, err
                }
                msg["annotations"] = &value
        }

        if obj.modified[netconf_profile_display_name] {
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

func (obj *NetconfProfile) UnmarshalJSON(body []byte) error {
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
                case "netconf_profile_is_default":
                        err = json.Unmarshal(value, &obj.netconf_profile_is_default)
                        if err == nil {
                                obj.valid[netconf_profile_netconf_profile_is_default] = true
                        }
                        break
                case "netconf_parameters":
                        err = json.Unmarshal(value, &obj.netconf_parameters)
                        if err == nil {
                                obj.valid[netconf_profile_netconf_parameters] = true
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid[netconf_profile_id_perms] = true
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid[netconf_profile_perms2] = true
                        }
                        break
                case "annotations":
                        err = json.Unmarshal(value, &obj.annotations)
                        if err == nil {
                                obj.valid[netconf_profile_annotations] = true
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid[netconf_profile_display_name] = true
                        }
                        break
                case "tag_refs":
                        err = json.Unmarshal(value, &obj.tag_refs)
                        if err == nil {
                                obj.valid[netconf_profile_tag_refs] = true
                        }
                        break
                case "telemetry_profile_back_refs":
                        err = json.Unmarshal(value, &obj.telemetry_profile_back_refs)
                        if err == nil {
                                obj.valid[netconf_profile_telemetry_profile_back_refs] = true
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *NetconfProfile) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified[netconf_profile_netconf_profile_is_default] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.netconf_profile_is_default)
                if err != nil {
                        return nil, err
                }
                msg["netconf_profile_is_default"] = &value
        }

        if obj.modified[netconf_profile_netconf_parameters] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.netconf_parameters)
                if err != nil {
                        return nil, err
                }
                msg["netconf_parameters"] = &value
        }

        if obj.modified[netconf_profile_id_perms] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified[netconf_profile_perms2] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified[netconf_profile_annotations] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.annotations)
                if err != nil {
                        return nil, err
                }
                msg["annotations"] = &value
        }

        if obj.modified[netconf_profile_display_name] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified[netconf_profile_tag_refs] {
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

func (obj *NetconfProfile) UpdateReferences() error {

        if obj.modified[netconf_profile_tag_refs] &&
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

func NetconfProfileByName(c contrail.ApiClient, fqn string) (*NetconfProfile, error) {
    obj, err := c.FindByName("netconf-profile", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*NetconfProfile), nil
}

func NetconfProfileByUuid(c contrail.ApiClient, uuid string) (*NetconfProfile, error) {
    obj, err := c.FindByUuid("netconf-profile", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*NetconfProfile), nil
}
