# SyntheticsSSLCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cipher** | Pointer to **string** |  | [optional] 
**Exponent** | Pointer to **float64** |  | [optional] 
**ExtKeyUsage** | Pointer to **[]string** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**Fingerprint256** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to [**SyntheticsSSLCertificateIssuer**](SyntheticsSSLCertificate_issuer.md) |  | [optional] 
**Modulus** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to [**SyntheticsSSLCertificateSubject**](SyntheticsSSLCertificate_subject.md) |  | [optional] 
**ValidFrom** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ValidTo** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### NewSyntheticsSSLCertificate

`func NewSyntheticsSSLCertificate() *SyntheticsSSLCertificate`

NewSyntheticsSSLCertificate instantiates a new SyntheticsSSLCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsSSLCertificateWithDefaults

`func NewSyntheticsSSLCertificateWithDefaults() *SyntheticsSSLCertificate`

NewSyntheticsSSLCertificateWithDefaults instantiates a new SyntheticsSSLCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCipher

`func (o *SyntheticsSSLCertificate) GetCipher() string`

GetCipher returns the Cipher field if non-nil, zero value otherwise.

### GetCipherOk

`func (o *SyntheticsSSLCertificate) GetCipherOk() (string, bool)`

GetCipherOk returns a tuple with the Cipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCipher

`func (o *SyntheticsSSLCertificate) HasCipher() bool`

HasCipher returns a boolean if a field has been set.

### SetCipher

`func (o *SyntheticsSSLCertificate) SetCipher(v string)`

SetCipher gets a reference to the given string and assigns it to the Cipher field.

### GetExponent

`func (o *SyntheticsSSLCertificate) GetExponent() float64`

GetExponent returns the Exponent field if non-nil, zero value otherwise.

### GetExponentOk

`func (o *SyntheticsSSLCertificate) GetExponentOk() (float64, bool)`

GetExponentOk returns a tuple with the Exponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExponent

`func (o *SyntheticsSSLCertificate) HasExponent() bool`

HasExponent returns a boolean if a field has been set.

### SetExponent

`func (o *SyntheticsSSLCertificate) SetExponent(v float64)`

SetExponent gets a reference to the given float64 and assigns it to the Exponent field.

### GetExtKeyUsage

`func (o *SyntheticsSSLCertificate) GetExtKeyUsage() []string`

GetExtKeyUsage returns the ExtKeyUsage field if non-nil, zero value otherwise.

### GetExtKeyUsageOk

`func (o *SyntheticsSSLCertificate) GetExtKeyUsageOk() ([]string, bool)`

GetExtKeyUsageOk returns a tuple with the ExtKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtKeyUsage

`func (o *SyntheticsSSLCertificate) HasExtKeyUsage() bool`

HasExtKeyUsage returns a boolean if a field has been set.

### SetExtKeyUsage

`func (o *SyntheticsSSLCertificate) SetExtKeyUsage(v []string)`

SetExtKeyUsage gets a reference to the given []string and assigns it to the ExtKeyUsage field.

### GetFingerprint

`func (o *SyntheticsSSLCertificate) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *SyntheticsSSLCertificate) GetFingerprintOk() (string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFingerprint

`func (o *SyntheticsSSLCertificate) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprint

`func (o *SyntheticsSSLCertificate) SetFingerprint(v string)`

SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.

### GetFingerprint256

`func (o *SyntheticsSSLCertificate) GetFingerprint256() string`

GetFingerprint256 returns the Fingerprint256 field if non-nil, zero value otherwise.

### GetFingerprint256Ok

`func (o *SyntheticsSSLCertificate) GetFingerprint256Ok() (string, bool)`

GetFingerprint256Ok returns a tuple with the Fingerprint256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFingerprint256

`func (o *SyntheticsSSLCertificate) HasFingerprint256() bool`

HasFingerprint256 returns a boolean if a field has been set.

### SetFingerprint256

`func (o *SyntheticsSSLCertificate) SetFingerprint256(v string)`

SetFingerprint256 gets a reference to the given string and assigns it to the Fingerprint256 field.

### GetIssuer

`func (o *SyntheticsSSLCertificate) GetIssuer() SyntheticsSSLCertificateIssuer`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SyntheticsSSLCertificate) GetIssuerOk() (SyntheticsSSLCertificateIssuer, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIssuer

`func (o *SyntheticsSSLCertificate) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuer

`func (o *SyntheticsSSLCertificate) SetIssuer(v SyntheticsSSLCertificateIssuer)`

SetIssuer gets a reference to the given SyntheticsSSLCertificateIssuer and assigns it to the Issuer field.

### GetModulus

`func (o *SyntheticsSSLCertificate) GetModulus() string`

GetModulus returns the Modulus field if non-nil, zero value otherwise.

### GetModulusOk

`func (o *SyntheticsSSLCertificate) GetModulusOk() (string, bool)`

GetModulusOk returns a tuple with the Modulus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModulus

`func (o *SyntheticsSSLCertificate) HasModulus() bool`

HasModulus returns a boolean if a field has been set.

### SetModulus

`func (o *SyntheticsSSLCertificate) SetModulus(v string)`

SetModulus gets a reference to the given string and assigns it to the Modulus field.

### GetProtocol

`func (o *SyntheticsSSLCertificate) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SyntheticsSSLCertificate) GetProtocolOk() (string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtocol

`func (o *SyntheticsSSLCertificate) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocol

`func (o *SyntheticsSSLCertificate) SetProtocol(v string)`

SetProtocol gets a reference to the given string and assigns it to the Protocol field.

### GetSerialNumber

`func (o *SyntheticsSSLCertificate) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *SyntheticsSSLCertificate) GetSerialNumberOk() (string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSerialNumber

`func (o *SyntheticsSSLCertificate) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumber

`func (o *SyntheticsSSLCertificate) SetSerialNumber(v string)`

SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.

### GetSubject

`func (o *SyntheticsSSLCertificate) GetSubject() SyntheticsSSLCertificateSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SyntheticsSSLCertificate) GetSubjectOk() (SyntheticsSSLCertificateSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubject

`func (o *SyntheticsSSLCertificate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubject

`func (o *SyntheticsSSLCertificate) SetSubject(v SyntheticsSSLCertificateSubject)`

SetSubject gets a reference to the given SyntheticsSSLCertificateSubject and assigns it to the Subject field.

### GetValidFrom

`func (o *SyntheticsSSLCertificate) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *SyntheticsSSLCertificate) GetValidFromOk() (time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValidFrom

`func (o *SyntheticsSSLCertificate) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### SetValidFrom

`func (o *SyntheticsSSLCertificate) SetValidFrom(v time.Time)`

SetValidFrom gets a reference to the given time.Time and assigns it to the ValidFrom field.

### GetValidTo

`func (o *SyntheticsSSLCertificate) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *SyntheticsSSLCertificate) GetValidToOk() (time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValidTo

`func (o *SyntheticsSSLCertificate) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.

### SetValidTo

`func (o *SyntheticsSSLCertificate) SetValidTo(v time.Time)`

SetValidTo gets a reference to the given time.Time and assigns it to the ValidTo field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

