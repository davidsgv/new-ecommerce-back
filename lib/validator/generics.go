package validator

// func min(dato interface{}, tagValue string) error {
// 	var mensaje string
// 	tagval, err := strconv.ParseInt(tagValue, 10, 64)
// 	if err != nil {
// 		return err
// 	}

// 	var validacion bool

// 	typeDato := reflect.TypeOf(dato)
// 	valueDato := reflect.ValueOf(dato)
// 	switch typeDato.Kind() {
// 	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
// 		d := valueDato.Uint()
// 		mensaje = fmt.Sprintf("%v debe ser al menos %s", dato, tagValue)
// 		validacion = d < uint64(tagval)
// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 		d := valueDato.Int()
// 		mensaje = fmt.Sprintf("%v debe ser al menos %s", dato, tagValue)
// 		validacion = d < tagval
// 	case reflect.Float32, reflect.Float64:
// 		d := valueDato.Float()
// 		mensaje = fmt.Sprintf("%v debe ser al menos %s", dato, tagValue)
// 		tagval, err := strconv.ParseFloat(tagValue, 64)
// 		if err != nil {
// 			return err
// 		}
// 		validacion = d < tagval
// 	case reflect.String:
// 		mensaje = fmt.Sprintf("%v debe tener una logitud de al menos %s", dato, tagValue)
// 		d := valueDato.String()
// 		validacion = len(d) < int(tagval)
// 	}

// 	if validacion {
// 		return errors.New(mensaje)
// 	}
// 	return nil
// }

// func max(dato interface{}, tagValue string) error {
// 	var mensaje string
// 	tagval, err := strconv.ParseInt(tagValue, 10, 64)
// 	if err != nil {
// 		return err
// 	}

// 	var validacion bool

// 	typeDato := reflect.TypeOf(dato)
// 	valueDato := reflect.ValueOf(dato)
// 	switch typeDato.Kind() {
// 	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
// 		d := valueDato.Uint()
// 		mensaje = fmt.Sprintf("%v no puede ser superior a %s", dato, tagValue)
// 		validacion = d > uint64(tagval)
// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 		d := valueDato.Int()
// 		mensaje = fmt.Sprintf("%v no puede ser superior a  %s", dato, tagValue)
// 		validacion = d > tagval
// 	case reflect.Float32, reflect.Float64:
// 		d := valueDato.Float()
// 		mensaje = fmt.Sprintf("%v no puede ser superior a %s", dato, tagValue)
// 		tagval, err := strconv.ParseFloat(tagValue, 64)
// 		if err != nil {
// 			return err
// 		}
// 		validacion = d > tagval
// 	case reflect.String:
// 		mensaje = fmt.Sprintf("%v debe tener una logitud inferior a %s", dato, tagValue)
// 		d := valueDato.String()
// 		validacion = len(d) > int(tagval)
// 	}

// 	if validacion {
// 		return errors.New(mensaje)
// 	}
// 	return nil
// }
