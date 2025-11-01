//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/ttsubo2000/contrail-go-api"
)

const (
	grpc_profile_grpc_profile_is_default = iota
	grpc_profile_grpc_parameters
	grpc_profile_id_perms
	grpc_profile_perms2
	grpc_profile_annotations
	grpc_profile_display_name
	grpc_profile_tag_refs
	grpc_profile_telemetry_profile_back_refs
	grpc_profile_max_
)

type GrpcProfile struct {
        contrail.ObjectBase
	grpc_profile_is_default bool
	grpc_parameters GrpcParameters
	id_perms IdPermsType
	perms2 PermType2
	annotations KeyValuePairs
	display_name string
	tag_refs contrail.ReferenceList
	telemetry_profile_back_refs contrail.ReferenceList
        valid [grpc_profile_max_] bool
        modified [grpc_profile_max_] bool
        baseMap map[string]contrail.ReferenceList
}

func (obj *GrpcProfile) GetType() string {
        return "grpc-profile"
}

func (obj *GrpcProfile) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project"}
        return name
}

func (obj *GrpcProfile) GetDefaultParentType() string {
        return "project"
}

func (obj *GrpcProfile) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *GrpcProfile) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *GrpcProfile) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *GrpcProfile) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *GrpcProfile) UpdateDone() {
        for i := range obj.modified { obj.modified[i] = false }
        obj.baseMap = nil
}


func (obj *GrpcProfile) GetGrpcProfileIsDefault() bool {
        return obj.grpc_profile_is_default
}

func (obj *GrpcProfile) SetGrpcProfileIsDefault(value bool) {
        obj.grpc_profile_is_default = value
        obj.modified[grpc_profile_grpc_profile_is_default] = true
}

func (obj *GrpcProfile) GetGrpcParameters() GrpcParameters {
        return obj.grpc_parameters
}

func (obj *GrpcProfile) SetGrpcParameters(value *GrpcParameters) {
        obj.grpc_parameters = *value
        obj.modified[grpc_profile_grpc_parameters] = true
}

func (obj *GrpcProfile) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *GrpcProfile) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified[grpc_profile_id_perms] = true
}

func (obj *GrpcProfile) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *GrpcProfile) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified[grpc_profile_perms2] = true
}

func (obj *GrpcProfile) GetAnnotations() KeyValuePairs {
        return obj.annotations
}

func (obj *GrpcProfile) SetAnnotations(value *KeyValuePairs) {
        obj.annotations = *value
        obj.modified[grpc_profile_annotations] = true
}

func (obj *GrpcProfile) GetDisplayName() string {
        return obj.display_name
}

func (obj *GrpcProfile) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified[grpc_profile_display_name] = true
}

func (obj *GrpcProfile) readTagRefs() error {
        if !obj.IsTransient() &&
                !obj.valid[grpc_profile_tag_refs] {
                err := obj.GetField(obj, "tag_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *GrpcProfile) GetTagRefs() (
        contrail.ReferenceList, error) {
        err := obj.readTagRefs()
        if err != nil {
                return nil, err
        }
        return obj.tag_refs, nil
}

func (obj *GrpcProfile) AddTag(
        rhs *Tag) error {
        err := obj.readTagRefs()
        if err != nil {
                return err
        }

        if !obj.modified[grpc_profile_tag_refs] {
                obj.storeReferenceBase("tag", obj.tag_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.tag_refs = append(obj.tag_refs, ref)
        obj.modified[grpc_profile_tag_refs] = true
        return nil
}

func (obj *GrpcProfile) DeleteTag(uuid string) error {
        err := obj.readTagRefs()
        if err != nil {
                return err
        }

        if !obj.modified[grpc_profile_tag_refs] {
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
        obj.modified[grpc_profile_tag_refs] = true
        return nil
}

func (obj *GrpcProfile) ClearTag() {
        if obj.valid[grpc_profile_tag_refs] &&
           !obj.modified[grpc_profile_tag_refs] {
                obj.storeReferenceBase("tag", obj.tag_refs)
        }
        obj.tag_refs = make([]contrail.Reference, 0)
        obj.valid[grpc_profile_tag_refs] = true
        obj.modified[grpc_profile_tag_refs] = true
}

func (obj *GrpcProfile) SetTagList(
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


func (obj *GrpcProfile) readTelemetryProfileBackRefs() error {
        if !obj.IsTransient() &&
                !obj.valid[grpc_profile_telemetry_profile_back_refs] {
                err := obj.GetField(obj, "telemetry_profile_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *GrpcProfile) GetTelemetryProfileBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readTelemetryProfileBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.telemetry_profile_back_refs, nil
}

func (obj *GrpcProfile) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified[grpc_profile_grpc_profile_is_default] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.grpc_profile_is_default)
                if err != nil {
                        return nil, err
                }
                msg["grpc_profile_is_default"] = &value
        }

        if obj.modified[grpc_profile_grpc_parameters] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.grpc_parameters)
                if err != nil {
                        return nil, err
                }
                msg["grpc_parameters"] = &value
        }

        if obj.modified[grpc_profile_id_perms] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified[grpc_profile_perms2] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified[grpc_profile_annotations] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.annotations)
                if err != nil {
                        return nil, err
                }
                msg["annotations"] = &value
        }

        if obj.modified[grpc_profile_display_name] {
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

func (obj *GrpcProfile) UnmarshalJSON(body []byte) error {
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
                case "grpc_profile_is_default":
                        err = json.Unmarshal(value, &obj.grpc_profile_is_default)
                        if err == nil {
                                obj.valid[grpc_profile_grpc_profile_is_default] = true
                        }
                        break
                case "grpc_parameters":
                        err = json.Unmarshal(value, &obj.grpc_parameters)
                        if err == nil {
                                obj.valid[grpc_profile_grpc_parameters] = true
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid[grpc_profile_id_perms] = true
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid[grpc_profile_perms2] = true
                        }
                        break
                case "annotations":
                        err = json.Unmarshal(value, &obj.annotations)
                        if err == nil {
                                obj.valid[grpc_profile_annotations] = true
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid[grpc_profile_display_name] = true
                        }
                        break
                case "tag_refs":
                        err = json.Unmarshal(value, &obj.tag_refs)
                        if err == nil {
                                obj.valid[grpc_profile_tag_refs] = true
                        }
                        break
                case "telemetry_profile_back_refs":
                        err = json.Unmarshal(value, &obj.telemetry_profile_back_refs)
                        if err == nil {
                                obj.valid[grpc_profile_telemetry_profile_back_refs] = true
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *GrpcProfile) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified[grpc_profile_grpc_profile_is_default] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.grpc_profile_is_default)
                if err != nil {
                        return nil, err
                }
                msg["grpc_profile_is_default"] = &value
        }

        if obj.modified[grpc_profile_grpc_parameters] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.grpc_parameters)
                if err != nil {
                        return nil, err
                }
                msg["grpc_parameters"] = &value
        }

        if obj.modified[grpc_profile_id_perms] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified[grpc_profile_perms2] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified[grpc_profile_annotations] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.annotations)
                if err != nil {
                        return nil, err
                }
                msg["annotations"] = &value
        }

        if obj.modified[grpc_profile_display_name] {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified[grpc_profile_tag_refs] {
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

func (obj *GrpcProfile) UpdateReferences() error {

        if obj.modified[grpc_profile_tag_refs] &&
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

func GrpcProfileByName(c contrail.ApiClient, fqn string) (*GrpcProfile, error) {
    obj, err := c.FindByName("grpc-profile", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*GrpcProfile), nil
}

func GrpcProfileByUuid(c contrail.ApiClient, uuid string) (*GrpcProfile, error) {
    obj, err := c.FindByUuid("grpc-profile", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*GrpcProfile), nil
}
