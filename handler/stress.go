package handler

import (
	"bufio"
	"crypto/sha512"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/iwalz/bench/config"
	stress "github.com/iwalz/bench/proto/stress"
	"golang.org/x/net/context"

	"database/sql"

	_ "github.com/Go-SQL-Driver/MySQL"
)

var conn *sql.DB

type Stress struct {
	config config.Config
}

func NewStressHandler(c config.Config) *Stress {
	return &Stress{
		config: c,
	}
}

// Create
func (s *Stress) StressCPU(ctx context.Context, req *stress.Request, rsp *stress.Response) error {
	for i := 0; i < s.config.Count; i++ {
		hash := sha512.New()
		b := []byte("This is a message that will get hashed " + strconv.Itoa(i))
		hash.Sum(b)
	}

	rsp.Message = "OK"
	return nil
}

func (s *Stress) StressMemory(ctx context.Context, req *stress.Request, rsp *stress.Response) error {
	data := make(map[int]string)
	// 1KB random text
	for i := 0; i < s.config.Count; i++ {
		data[i] = "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" + strconv.Itoa(i)
	}
	rsp.Message = "OK"
	return nil
}

func (s *Stress) StressIO(ctx context.Context, req *stress.Request, rsp *stress.Response) error {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	name := strconv.Itoa(r.Intn(1000))
	f, _ := os.Create("/tmp/" + name)
	w := bufio.NewWriter(f)

	// 1KB random text
	for i := 0; i < s.config.Count/10; i++ {
		data := "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" + strconv.Itoa(i)
		w.WriteString(data)
		w.Flush()
	}

	f.Close()
	os.Remove("/tmp/" + name)
	rsp.Message = "OK"
	return nil
}

func (s *Stress) StressNetwork(ctx context.Context, req *stress.NetworkRequest, rsp *stress.NetworkResponse) error {
	msg := ""

	for i := 0; i < s.config.Count/10; i++ {
		msg = msg + req.Message
	}

	rsp.Message = msg
	return nil
}

func (s *Stress) StressRDS(ctx context.Context, req *stress.Request, rsp *stress.Response) error {
	db, err := sql.Open("mysql", s.config.User+":"+s.config.Password+"@tcp("+s.config.Endpoint+")/"+s.config.Database+"?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	for i := 0; i < s.config.Count/100; i++ {
		stmt, err := db.Prepare("INSERT test SET foo=?,bar=?")
		if err != nil {
			fmt.Println(err)
			return err
		}
		stmt.Exec("abcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghi", "abcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghi")
	}

	stmt, _ := db.Prepare("DELETE FROM test")
	stmt.Exec()
	return nil
}

func (s *Stress) StressRDSIOPs(ctx context.Context, req *stress.Request, rsp *stress.Response) error {
	if conn == nil {
		var err error
		conn, err = sql.Open("mysql", s.config.User+":"+s.config.Password+"@tcp("+s.config.Endpoint+")/"+s.config.Database+"?charset=utf8")
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	for i := 0; i < s.config.Count/100; i++ {
		stmt, err := conn.Prepare("INSERT test SET foo=?,bar=?")
		if err != nil {
			fmt.Println(err)
			return err
		}
		stmt.Exec("abcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghi", "abcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghi")
	}

	stmt, _ := conn.Prepare("DELETE FROM test")
	stmt.Exec()
	return nil
}
