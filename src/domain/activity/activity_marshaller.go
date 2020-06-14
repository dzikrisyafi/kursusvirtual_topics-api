package activity

type PublicActivity struct {
	ID int `json:"id"`
}

func (activitys Activitys) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(activitys))
	for index, activity := range activitys {
		result[index] = activity.Marshall(isPublic)
	}

	return result
}

func (activity *Activity) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicActivity{
			ID: activity.ID,
		}
	}

	return activity
}
