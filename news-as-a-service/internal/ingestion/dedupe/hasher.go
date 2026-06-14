package dedup

import (
	"crypto/sha256"
	"encoding/hex"
)

func URLHash(
	url string,
) string {

	sum :=
		sha256.Sum256(
			[]byte(url),
		)

	return hex.EncodeToString(
		sum[:],
	)
}

func EntityDedupeKey(
	source string,
	title string,
	date string,
) string {

	payload :=
		source +
			title +
			date

	sum :=
		sha256.Sum256(
			[]byte(payload),
		)

	return hex.EncodeToString(
		sum[:],
	)
}
