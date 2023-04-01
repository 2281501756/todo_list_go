package serializer

import "todo_list/model"

type task = struct {
	model.Task
	View uint64 `json:"view"`
}

func BuildTask(t *model.Task) (res task) {
	res.ID = t.ID
	res.Title = t.Title
	res.UserId = t.UserId
	res.User = t.User
	res.Content = t.Content
	res.Status = t.Status
	res.StartTime = t.StartTime
	res.EndTime = t.EndTime
	res.CreatedAt = t.CreatedAt
	res.UpdatedAt = t.UpdatedAt
	res.DeletedAt = t.DeletedAt
	res.View = t.GetView()
	return
}
func BuildTasks(t []model.Task) (res []task) {
	for _, i := range t {
		res = append(res, BuildTask(&i))
	}
	return
}
