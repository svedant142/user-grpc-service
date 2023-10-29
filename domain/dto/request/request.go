package request

import "errors"

func ValidateUserID(userID int64) error {
	if userID <= 0 {
		return errors.New("ID cannot be less than 1")
	}
	return nil
}

func ValidateUserIDs(userIDs []int64) error {
	if userIDs == nil {
		return errors.New("no ids in request")
	}
	repeatedIDs := make(map[int64]bool)
	for _, id := range userIDs {
		if repeatedIDs[id] {
			return errors.New("repeated IDs in request")
		}
		if err := ValidateUserID(id); err != nil {
			return err
		}
		repeatedIDs[id]=true
	}
	return nil
}
