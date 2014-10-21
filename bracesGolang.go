func (d *digest) Write(p []byte) (nn int, err error) { // HL
	nn = len(p)
	d.len += uint64(nn)
	if d.nx > 0 { // HL
		n := copy(d.x[d.nx:], p)
		d.nx += n
		if d.nx == chunk { // HL
			block(d, d.x[:])
			d.nx = 0
		} // HL
		p = p[n:]
	} // HL
	if len(p) >= chunk { // HL
		n := len(p) &^ (chunk - 1)
		block(d, p[:n])
		p = p[n:]
	} // HL
	if len(p) > 0 { // HL
		d.nx = copy(d.x[:], p)
	} // HL
	return
} // HL