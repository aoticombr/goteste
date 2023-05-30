package lib

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var logo = `
 
╔═╗╔═╗╔╦╗╦  ╔═╗╔═╗  ╔═╗╔═╗╦
╠═╣║ ║ ║ ║  ║ ╦║ ║  ╠═╣╠═╝║
╩ ╩╚═╝ ╩ ╩  ╚═╝╚═╝  ╩ ╩╩  ╩
`

var (
	keyFile     string
	path_log    string
	path_xml    string
	path_sftp   string
	sftp_ip     string
	sftp_user   string
	ora_ip      string
	ora_port    string
	ora_user    string
	ora_pass    string
	ora_schema  string
	pg_ip       string
	pg_port     string
	pg_user     string
	pg_pass     string
	pg_database string
	port        string
	secret_key  string
	debug       bool
	view        bool
)

func Setsecret_key(value string) {
	secret_key = value
}
func GetSecret_key() string {
	return secret_key
}

func Setpg_ip(value string) {
	pg_ip = value
}
func Setpg_port(value string) {
	pg_port = value
}
func Setpg_user(value string) {
	pg_user = value
}
func Setpg_pass(value string) {
	pg_pass = value
}
func Setpg_database(value string) {
	pg_database = value
}

func Getpg_ip() string {
	return pg_ip
}
func Getpg_port() string {
	return pg_port
}
func Getpg_user() string {
	return pg_user
}
func Getpg_pass() string {
	return pg_pass
}
func Getpg_database() string {
	return pg_database
}

func GetPathLOG() string {
	return path_log
}
func GetPathXML() string {
	return path_xml
}
func GetPathSFTP() string {
	return path_sftp
}
func GetLog() string {
	return logo
}
func GetKeyFile() string {
	return keyFile
}

func GetSFTP_IP() string {
	return sftp_ip
}
func GetSFTP_USER() string {
	return sftp_user
}

func GetOra_ip() string {
	return ora_ip
}
func GetOra_port() string {
	return ora_port
}
func GetOra_portInt() int {
	intValue, _ := strconv.Atoi(GetOra_port())
	return intValue
}
func GetPort() string {
	return port
}
func GetOra_user() string {
	return ora_user
}
func GetOra_pass() string {
	return ora_pass
}
func GetOra_schema() string {
	return ora_schema
}
func Getdebug() bool {
	return debug
}
func Getview() bool {
	return view
}

func SetkeyFile(value string) {
	keyFile = value
}
func Setpath_log(value string) {
	path_log = value
}

func Setpath_xml(value string) {
	path_xml = value
}
func Setpath_sftp(value string) {
	path_sftp = value
}
func Setsftp_ip(value string) {
	sftp_ip = value
}
func Setsftp_user(value string) {
	sftp_user = value
}
func Setora_ip(value string) {
	ora_ip = value
}
func Setora_port(value string) {
	ora_port = value
}

func SetPort(value string) {
	port = value
}
func Setora_user(value string) {
	ora_user = value
}
func Setora_pass(value string) {
	ora_pass = value
}
func Setora_schema(value string) {
	ora_schema = value
}
func Setdebug(value string) error {
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	debug = boolValue
	return nil
}
func Setview(value string) error {
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	view = boolValue
	return nil
}

func LerEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	SetkeyFile(os.Getenv("keyFile"))
	Setpath_log(os.Getenv("path_log"))
	Setpath_xml(os.Getenv("path_xml"))
	Setpath_sftp(os.Getenv("path_sftp"))
	Setsftp_ip(os.Getenv("sftp_ip"))
	Setsftp_user(os.Getenv("sftp_user"))
	Setora_ip(os.Getenv("ora_ip"))

	Setora_user(os.Getenv("ora_user"))
	Setora_pass(os.Getenv("ora_pass"))
	Setora_schema(os.Getenv("ora_schema"))
	Setora_port(os.Getenv("ora_port"))

	Setpg_database(os.Getenv("pg_database"))
	Setpg_ip(os.Getenv("pg_ip"))
	Setpg_pass(os.Getenv("pg_pass"))
	Setpg_user(os.Getenv("pg_user"))
	Setpg_port(os.Getenv("pg_port"))

	Setsecret_key(os.Getenv("secret_key"))

	SetPort(os.Getenv("port"))
	err = Setdebug(os.Getenv("debug"))
	if err != nil {
		return err
	}
	err = Setview(os.Getenv("view"))
	if err != nil {
		return err
	}

	return nil
}
