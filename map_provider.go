package cfg 

// MapProvider provides a simple implementation of the provider whereby it 
// just returns a stored map 
type MapProvider struct{
    Map map[string]string
}

// Provide implements the provider interface
func  (mp MapProvider)Provide()(map[string]string, error) {
    return mp.Map, nil 
}