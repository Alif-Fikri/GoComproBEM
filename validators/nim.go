package validators

import "github.com/go-playground/validator/v10"

func ValidateNIM(fl validator.FieldLevel) bool {
    nim := fl.Field().String()
    
    if len(nim) != 11 {
        return false
    }
    
    if nim[:3] != "415" && nim[:3] != "418" {
        return false
    }
    
    for _, c := range nim {
        if c < '0' || c > '9' {
            return false
        }
    }
    
    return true
}