// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/testifysec/archivist/ent/attestation"
	"github.com/testifysec/archivist/ent/attestationcollection"
	"github.com/testifysec/archivist/ent/dsse"
	"github.com/testifysec/archivist/ent/payloaddigest"
	"github.com/testifysec/archivist/ent/schema"
	"github.com/testifysec/archivist/ent/signature"
	"github.com/testifysec/archivist/ent/statement"
	"github.com/testifysec/archivist/ent/subject"
	"github.com/testifysec/archivist/ent/subjectdigest"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	attestationFields := schema.Attestation{}.Fields()
	_ = attestationFields
	// attestationDescType is the schema descriptor for type field.
	attestationDescType := attestationFields[0].Descriptor()
	// attestation.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	attestation.TypeValidator = attestationDescType.Validators[0].(func(string) error)
	attestationcollectionFields := schema.AttestationCollection{}.Fields()
	_ = attestationcollectionFields
	// attestationcollectionDescName is the schema descriptor for name field.
	attestationcollectionDescName := attestationcollectionFields[0].Descriptor()
	// attestationcollection.NameValidator is a validator for the "name" field. It is called by the builders before save.
	attestationcollection.NameValidator = attestationcollectionDescName.Validators[0].(func(string) error)
	dsseFields := schema.Dsse{}.Fields()
	_ = dsseFields
	// dsseDescGitbomSha256 is the schema descriptor for gitbom_sha256 field.
	dsseDescGitbomSha256 := dsseFields[0].Descriptor()
	// dsse.GitbomSha256Validator is a validator for the "gitbom_sha256" field. It is called by the builders before save.
	dsse.GitbomSha256Validator = dsseDescGitbomSha256.Validators[0].(func(string) error)
	// dsseDescPayloadType is the schema descriptor for payload_type field.
	dsseDescPayloadType := dsseFields[1].Descriptor()
	// dsse.PayloadTypeValidator is a validator for the "payload_type" field. It is called by the builders before save.
	dsse.PayloadTypeValidator = dsseDescPayloadType.Validators[0].(func(string) error)
	payloaddigestFields := schema.PayloadDigest{}.Fields()
	_ = payloaddigestFields
	// payloaddigestDescAlgorithm is the schema descriptor for algorithm field.
	payloaddigestDescAlgorithm := payloaddigestFields[0].Descriptor()
	// payloaddigest.AlgorithmValidator is a validator for the "algorithm" field. It is called by the builders before save.
	payloaddigest.AlgorithmValidator = payloaddigestDescAlgorithm.Validators[0].(func(string) error)
	// payloaddigestDescValue is the schema descriptor for value field.
	payloaddigestDescValue := payloaddigestFields[1].Descriptor()
	// payloaddigest.ValueValidator is a validator for the "value" field. It is called by the builders before save.
	payloaddigest.ValueValidator = payloaddigestDescValue.Validators[0].(func(string) error)
	signatureFields := schema.Signature{}.Fields()
	_ = signatureFields
	// signatureDescKeyID is the schema descriptor for key_id field.
	signatureDescKeyID := signatureFields[0].Descriptor()
	// signature.KeyIDValidator is a validator for the "key_id" field. It is called by the builders before save.
	signature.KeyIDValidator = signatureDescKeyID.Validators[0].(func(string) error)
	// signatureDescSignature is the schema descriptor for signature field.
	signatureDescSignature := signatureFields[1].Descriptor()
	// signature.SignatureValidator is a validator for the "signature" field. It is called by the builders before save.
	signature.SignatureValidator = signatureDescSignature.Validators[0].(func(string) error)
	statementFields := schema.Statement{}.Fields()
	_ = statementFields
	// statementDescPredicate is the schema descriptor for predicate field.
	statementDescPredicate := statementFields[0].Descriptor()
	// statement.PredicateValidator is a validator for the "predicate" field. It is called by the builders before save.
	statement.PredicateValidator = statementDescPredicate.Validators[0].(func(string) error)
	subjectFields := schema.Subject{}.Fields()
	_ = subjectFields
	// subjectDescName is the schema descriptor for name field.
	subjectDescName := subjectFields[0].Descriptor()
	// subject.NameValidator is a validator for the "name" field. It is called by the builders before save.
	subject.NameValidator = subjectDescName.Validators[0].(func(string) error)
	subjectdigestFields := schema.SubjectDigest{}.Fields()
	_ = subjectdigestFields
	// subjectdigestDescAlgorithm is the schema descriptor for algorithm field.
	subjectdigestDescAlgorithm := subjectdigestFields[0].Descriptor()
	// subjectdigest.AlgorithmValidator is a validator for the "algorithm" field. It is called by the builders before save.
	subjectdigest.AlgorithmValidator = subjectdigestDescAlgorithm.Validators[0].(func(string) error)
	// subjectdigestDescValue is the schema descriptor for value field.
	subjectdigestDescValue := subjectdigestFields[1].Descriptor()
	// subjectdigest.ValueValidator is a validator for the "value" field. It is called by the builders before save.
	subjectdigest.ValueValidator = subjectdigestDescValue.Validators[0].(func(string) error)
}
