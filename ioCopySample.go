//STARTIO OMIT
func Copy(dst Writer, src Reader) (written int64, err error) {
	buf := make([]byte, 32*1024)
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if ew != nil {
				err = ew
				break
			}
		}
		if er == EOF {
			break
		}
		if er != nil {
			err = er
			break
		}
	}
	return written, err
}
//ENDIO OMIT
//STARTAPP OMIT
import (
	"io"
	"sql"
)

func packageDBCopy() error {
	db, err := sql.Open(/* ... */)
	if err != nil {
		// Alternatively: retry, log, other library-specific recovery actions
		return err
	}
	// ...
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	// ...
	return nil
}
//ENDAPP OMIT
//STARTAPPCLOSE OMIT
import (
	"io"
	"sql"
)

func packageDBCopy() error {
	db, err := sql.Open(/* ... */)
	if err != nil {
		return err
	}
	// ...
	_, err = io.Copy(dst, src)
	if err != nil {
		db.Close() // HL
		return err
	}
	// ...
	db.Close() // HL
	return nil
}
//ENDAPPCLOSE OMIT
//STARTAPPDEFER OMIT
import (
	"io"
	"sql"
)

func packageDBCopy() error {
	db, err := sql.Open(/* ... */)
	if err != nil {
		return err
	}
	defer db.Close() // HL
	// ...
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	// ...
	return nil
}
//ENDAPPDEFER OMIT