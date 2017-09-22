package aliases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_User_Validate(t *testing.T) {
	t.Parallel()

	someString := StringType("aaaaaaaaaaaaaaaaaaaab")
	validUser := User{
		FirstName:     "firstName",
		LastName:      "lastName",
		FamilyMembers: 99,
		SomeFloat:     5.55,
		SomeMap: map[string]int{
			"test":  1,
			"test2": 2,
		},
		SomePointer: &someString,
	}

	t.Run("valid", func(t *testing.T) {
		assert.NoError(t, validUser.Validate())
	})

	t.Run("invalid", func(t *testing.T) {

		t.Run("first_name: too short, using overridden alias validator", func(t *testing.T) {
			r := validUser
			r.FirstName = "aa"

			err := r.Validate()
			require.NoError(t, err)
		})

		t.Run("first_name: too short, using alias validator", func(t *testing.T) {
			r := validUser
			r.FirstName = "aa"

			err := r.FirstName.Validate()
			require.NotNil(t, err)
			assert.Equal(t, `shorter than 3 chars`, err.Error())
		})

		t.Run("some_pointer: not_null rule", func(t *testing.T) {
			r := validUser
			r.SomePointer = nil

			err := r.Validate()
			require.NotNil(t, err)
			assert.Equal(t, `[SomePointer: cannot be nil]`, err.Error())
		})

		t.Run("some_pointer: using overridden alias validator", func(t *testing.T) {
			r := validUser
			someString := StringType("aaa")
			r.SomePointer = &someString

			err := r.Validate()
			require.NotNil(t, err)
			assert.Equal(t, `[SomePointer: shorter than 20 chars]`, err.Error())
		})
	})
}
