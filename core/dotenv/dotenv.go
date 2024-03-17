package dotenv

import (
	"app/core/utils"
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const doubleQuoteSpecialChars = "\\\n\r\"!$`"

type DotEnv struct {
}

func (o DotEnv) GetBool(key string, defaultVal string, errMsgRequired string) bool {
	required := errMsgRequired != ""
	s := o.GetEnv(key)

	if !required && s.Value == "" {
		s.Value = defaultVal
	}

	var ok bool

	ret := s.ToBool(&ok)

	if !ok && required {
		log.Fatal(errMsgRequired)
	}

	return ret
}

func (o DotEnv) GetInt16(key string, defaultVal string, errMsgRequired string) int16 {
	required := errMsgRequired != ""
	s := o.GetEnv(key)

	if !required && s.Value == "" {
		s.Value = defaultVal
	}

	var ok bool

	ret := s.ToInt32(&ok)

	if !ok && required {
		log.Fatal(errMsgRequired)
	}

	return int16(ret)
}

func (o DotEnv) GetUint16(key string, defaultVal string, errMsgRequired string) uint16 {
	required := errMsgRequired != ""
	s := o.GetEnv(key)

	if !required && s.Value == "" {
		s.Value = defaultVal
	}

	var ok bool

	ret := s.ToUint32(&ok)

	if !ok && required {
		log.Fatal(errMsgRequired)
	}

	return uint16(ret)
}

func (o DotEnv) GetInt(key string, defaultVal string, errMsgRequired string) int32 {
	required := errMsgRequired != ""
	s := o.GetEnv(key)

	if !required && s.Value == "" {
		s.Value = defaultVal
	}

	var ok bool

	ret := s.ToInt32(&ok)

	if !ok && required {
		log.Fatal(errMsgRequired)
	}

	return ret
}

func (o DotEnv) GetUint(key string, defaultVal string, errMsgRequired string) uint32 {
	required := errMsgRequired != ""
	s := o.GetEnv(key)

	if !required && s.Value == "" {
		s.Value = defaultVal
	}

	var ok bool

	ret := s.ToUint32(&ok)

	if !ok && required {
		log.Fatal(errMsgRequired)
	}

	return ret
}

func (o DotEnv) GetInt64(key string, defaultVal string, errMsgRequired string) int64 {
	required := errMsgRequired != ""
	s := o.GetEnv(key)

	if !required && s.Value == "" {
		s.Value = defaultVal
	}

	var ok bool

	ret := s.ToInt64(&ok)

	if !ok && required {
		log.Fatal(errMsgRequired)
	}

	return ret
}

func (o DotEnv) GetUint64(key string, defaultVal string, errMsgRequired string) uint64 {
	required := errMsgRequired != ""
	s := o.GetEnv(key)

	if !required && s.Value == "" {
		s.Value = defaultVal
	}

	var ok bool

	ret := s.ToUint64(&ok)

	if !ok && required {
		log.Fatal(errMsgRequired)
	}

	return ret
}

func (o DotEnv) GetString(key string, defaultVal string, errMsgRequired string) string {
	required := errMsgRequired != ""
	s := o.GetEnv(key)

	if !required && s.Value == "" {
		s.Value = defaultVal
	}

	if s.Value == "" && required {
		log.Fatal(errMsgRequired)
	}

	return s.Value
}

func (o DotEnv) GetStringArray(key string, sep string, errMsgRequired string) []string {
	required := errMsgRequired != ""
	s := o.GetEnv(key)

	if s.Value == "" && required {
		log.Fatal(errMsgRequired)
	}

	items := utils.UtilString{}.Split(s.Value, sep)

	return items
}

func (o DotEnv) GetFloat(key string, defaultVal string, errMsgRequired string) float64 {
	required := errMsgRequired != ""
	s := o.GetEnv(key)

	if !required && s.Value == "" {
		s.Value = defaultVal
	}

	var ok bool

	ret := s.ToFloat(&ok)

	if !ok && required {
		log.Fatal(errMsgRequired)
	}

	return ret
}

func (obj DotEnv) GetEnv(key string) utils.UtilString {
	return utils.UtilString{Value: os.Getenv(key)}
}

// Load will read your env file(s) and load them into ENV for this process.
//
// Call this function as close as possible to the start of your program (ideally in main)
//
// If you call Load without any args it will default to loading .env in the current path
//
// You can otherwise tell it which files to load (there can be more than one) like
//
//		godotenv.Load("fileone", "filetwo")
//
// It's important to note that it WILL NOT OVERRIDE an env variable that already exists - consider the .env file to set dev vars or sensible defaults
func (obj DotEnv) Load(filenames ...string) (err error) {
	filenames = obj.filenamesOrDefault(filenames)

	for _, filename := range filenames {
		err = obj.loadFile(filename, false)
		if err != nil {
			return // return early on a spazout
		}
	}
	return
}

// Overload will read your env file(s) and load them into ENV for this process.
//
// Call this function as close as possible to the start of your program (ideally in main)
//
// If you call Overload without any args it will default to loading .env in the current path
//
// You can otherwise tell it which files to load (there can be more than one) like
//
//		godotenv.Overload("fileone", "filetwo")
//
// It's important to note this WILL OVERRIDE an env variable that already exists - consider the .env file to forcefilly set all vars.
func (obj DotEnv) Overload(filenames ...string) (err error) {
	filenames = obj.filenamesOrDefault(filenames)

	for _, filename := range filenames {
		err = obj.loadFile(filename, true)
		if err != nil {
			return // return early on a spazout
		}
	}
	return
}

// Read all env (with same file loading semantics as Load) but return values as
// a map rather than automatically writing values into env
func (obj DotEnv) Read(filenames ...string) (envMap map[string]string, err error) {
	filenames = obj.filenamesOrDefault(filenames)
	envMap = make(map[string]string)

	for _, filename := range filenames {
		individualEnvMap, individualErr := obj.readFile(filename)

		if individualErr != nil {
			err = individualErr
			return // return early on a spazout
		}

		for key, value := range individualEnvMap {
			envMap[key] = value
		}
	}

	return
}

// Parse reads an env file from io.Reader, returning a map of keys and values.
func (obj DotEnv) Parse(r io.Reader) (envMap map[string]string, err error) {
	envMap = make(map[string]string)

	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return
	}

	for _, fullLine := range lines {
		if !obj.isIgnoredLine(fullLine) {
			var key, value string
			key, value, err = obj.parseLine(fullLine, envMap)

			if err != nil {
				return
			}
			envMap[key] = value
		}
	}
	return
}

//Unmarshal reads an env file from a string, returning a map of keys and values.
func (obj DotEnv) Unmarshal(str string) (envMap map[string]string, err error) {
	return obj.Parse(strings.NewReader(str))
}

// Exec loads env vars from the specified filenames (empty map falls back to default)
// then executes the cmd specified.
//
// Simply hooks up os.Stdin/err/out to the command and calls Run()
//
// If you want more fine grained control over your command it's recommended
// that you use `Load()` or `Read()` and the `os/exec` package yourself.
func (obj DotEnv) Exec(filenames []string, cmd string, cmdArgs []string) error {
	obj.Load(filenames...)

	command := exec.Command(cmd, cmdArgs...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}

// Write serializes the given environment and writes it to a file
func (obj DotEnv) Write(envMap map[string]string, filename string) error {
	content, err := obj.Marshal(envMap)
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content + "\n")
	if err != nil {
		return err
	}
	file.Sync()
	return err
}

// Marshal outputs the given environment as a dotenv-formatted environment file.
// Each line is in the format: KEY="VALUE" where VALUE is backslash-escaped.
func (obj DotEnv) Marshal(envMap map[string]string) (string, error) {
	lines := make([]string, 0, len(envMap))
	for k, v := range envMap {
		if d, err := strconv.Atoi(v); err == nil {
			lines = append(lines, fmt.Sprintf(`%s=%d`, k, d))
		} else {
			lines = append(lines, fmt.Sprintf(`%s="%s"`, k, obj.doubleQuoteEscape(v)))
		}
	}
	sort.Strings(lines)
	return strings.Join(lines, "\n"), nil
}

func (obj DotEnv) filenamesOrDefault(filenames []string) []string {
	if len(filenames) == 0 {
		return []string{".env"}
	}
	return filenames
}

func (obj DotEnv) loadFile(filename string, overload bool) error {
	envMap, err := obj.readFile(filename)
	if err != nil {
		return err
	}

	currentEnv := map[string]bool{}
	rawEnv := os.Environ()
	for _, rawEnvLine := range rawEnv {
		key := strings.Split(rawEnvLine, "=")[0]
		currentEnv[key] = true
	}

	for key, value := range envMap {
		if !currentEnv[key] || overload {
			os.Setenv(key, value)
		}
	}

	return nil
}

func (obj DotEnv) readFile(filename string) (envMap map[string]string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	return obj.Parse(file)
}

func (obj DotEnv) parseLine(line string, envMap map[string]string) (key string, value string, err error) {
	var exportRegex = regexp.MustCompile(`^\s*(?:export\s+)?(.*?)\s*$`)

	if len(line) == 0 {
		err = errors.New("zero length string")
		return
	}

	// ditch the comments (but keep quoted hashes)
	if strings.Contains(line, "#") {
		segmentsBetweenHashes := strings.Split(line, "#")
		var segmentsToKeep []string
		for _, segment := range segmentsBetweenHashes {
			segmentsToKeep = append(segmentsToKeep, segment)
			if segment[len(segment)-1] == ' ' {
				break
			}
		}
		line = strings.Join(segmentsToKeep, "#")
	}

	firstEquals := strings.Index(line, "=")
	firstColon := strings.Index(line, ":")
	splitString := strings.SplitN(line, "=", 2)
	if firstColon != -1 && (firstColon < firstEquals || firstEquals == -1) {
		//this is a yaml-style line
		splitString = strings.SplitN(line, ":", 2)
	}

	if len(splitString) != 2 {
		err = errors.New("can't separate key from value")
		return
	}

	// Parse the key
	/*key = splitString[0]
	if strings.HasPrefix(key, "export") {
		key = strings.TrimPrefix(key, "export")
	}
	key = strings.TrimSpace(key)*/

	key = exportRegex.ReplaceAllString(splitString[0], "$1")

	// Parse the value
	value = obj.parseValue(splitString[1], envMap)
	return
}

func (obj DotEnv) parseValue(value string, envMap map[string]string) string {

	var (
		singleQuotesRegex  = regexp.MustCompile(`\A'(.*)'\z`)
		doubleQuotesRegex  = regexp.MustCompile(`\A"(.*)"\z`)
		escapeRegex        = regexp.MustCompile(`\\.`)
		unescapeCharsRegex = regexp.MustCompile(`\\([^$])`)
	)

	// trim
	value = strings.Trim(value, " ")

	// check if we've got quoted values or possible escapes
	if len(value) > 1 {
		singleQuotes := singleQuotesRegex.FindStringSubmatch(value)

		doubleQuotes := doubleQuotesRegex.FindStringSubmatch(value)

		if singleQuotes != nil || doubleQuotes != nil {
			// pull the quotes off the edges
			value = value[1 : len(value)-1]
		}

		if doubleQuotes != nil {
			// expand newlines
			value = escapeRegex.ReplaceAllStringFunc(value, func(match string) string {
				c := strings.TrimPrefix(match, `\`)
				switch c {
				case "n":
					return "\n"
				case "r":
					return "\r"
				default:
					return match
				}
			})
			// unescape characters
			value = unescapeCharsRegex.ReplaceAllString(value, "$1")
		}

		if singleQuotes == nil {
			value = obj.expandVariables(value, envMap)
		}
	}

	return value
}

func (obj DotEnv) expandVariables(v string, m map[string]string) string {
	var expandVarRegex = regexp.MustCompile(`(\\)?(\$)(\()?\{?([A-Z0-9_]+)?\}?`)

	return expandVarRegex.ReplaceAllStringFunc(v, func(s string) string {
		submatch := expandVarRegex.FindStringSubmatch(s)

		if submatch == nil {
			return s
		}
		if submatch[1] == "\\" || submatch[2] == "(" {
			return submatch[0][1:]
		} else if submatch[4] != "" {
			return m[submatch[4]]
		}
		return s
	})
}

func (obj DotEnv) isIgnoredLine(line string) bool {
	trimmedLine := strings.TrimSpace(line)
	return len(trimmedLine) == 0 || strings.HasPrefix(trimmedLine, "#")
}

func (obj DotEnv) doubleQuoteEscape(line string) string {
	for _, c := range doubleQuoteSpecialChars {
		toReplace := "\\" + string(c)
		if c == '\n' {
			toReplace = `\n`
		}
		if c == '\r' {
			toReplace = `\r`
		}
		line = strings.Replace(line, string(c), toReplace, -1)
	}
	return line
}
