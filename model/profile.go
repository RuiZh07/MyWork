package model
import(
	"database/sql"
)

type ProfileMenu struct {
    profileLink string
	ShowCreateProfileButton bool
	ProfilePages            []string
}

type ProfileData struct{
	ProfileName string
	ProfileLinks []string
}

type Profile struct {
    ProfileID    int    `db:"profile_id"`
    UserID       int    `db:"user_id"`
    UserEmail    string `db:"user_email"`
    Name         string `db:"name"`
    Activation   bool   `db:"activation"`
    Link1        sql.NullString `db:"link1"`
    Link2        sql.NullString `db:"link2"`
    Link3        sql.NullString `db:"link3"`
    Link4        sql.NullString `db:"link4"`
    Link5        sql.NullString `db:"link5"`
    Link6        sql.NullString `db:"link6"`
    Link7        sql.NullString `db:"link7"`
    Link8        sql.NullString `db:"link8"`
    Link9        sql.NullString `db:"link9"`
    Link10       sql.NullString `db:"link10"`
}

type ProfileForNFC struct {
    UserName string
    UserProfilePicture string
    UserUniversity string
    ProfileLinks []string
}