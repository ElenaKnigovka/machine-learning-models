
func ReadFileContents(filename string) (string, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

