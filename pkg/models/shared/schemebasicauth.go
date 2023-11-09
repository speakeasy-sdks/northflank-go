// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SchemeBasicAuth struct {
	Password string `security:"name=password"`
	Username string `security:"name=username"`
}

func (o *SchemeBasicAuth) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *SchemeBasicAuth) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
