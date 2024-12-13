-- name: InsertSong :one
INSERT INTO song (song, group_name, release_date, text, link)
values (@song, @group_name, @release_date, @text, @link)
RETURNING id;

-- name: DeleteSong :exec
DELETE FROM song
WHERE id = @id;

-- name: UpdateSong :exec
UPDATE song
SET song = @song,
    group_name = @group_name,
    release_date = @release_date,
    text = @text,
    link = @link
WHERE id = @id;

-- name: SelectSongText :one
SELECT text
FROM song
WHERE id = @id;