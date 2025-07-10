package helpers

import "triple-s/models"

func NtoB(notes [][]string) models.Buckets {
	var buckets models.Buckets
	for _, note := range notes {
		bucket := models.Bucket{
			Name:             note[0],
			CreationTime:     note[1],
			LastModifiedTime: note[2],
			Status:           note[3],
		}
		buckets.Buckets = append(buckets.Buckets, bucket)
	}
	return buckets
}

func BtoN(buckets models.Buckets) [][]string {
	var notes [][]string
	for _, bucket := range buckets.Buckets {
		note := []string{
			bucket.Name,
			bucket.CreationTime,
			bucket.LastModifiedTime,
			bucket.Status,
		}
		notes = append(notes, note)
	}
	return notes
}

func NtoO(notes [][]string) models.Objects {
	var objects models.Objects
	for _, note := range notes {
		object := models.Object{
			ObjectKey:    note[0],
			ContentType:  note[1],
			Size:         note[2],
			LastModified: note[3],
		}
		objects.Objects = append(objects.Objects, object)
	}
	return objects
}

func OtoN(objects models.Objects) [][]string {
	var notes [][]string
	for _, object := range objects.Objects {
		note := []string{
			object.ObjectKey,
			object.ContentType,
			object.Size,
			object.LastModified,
		}
		notes = append(notes, note)
	}
	return notes
}
