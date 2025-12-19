package auth

type Clearance int
type Action int
type Content int

const (
	// ClearanceClassified users can comment, edit their own comments, delete their own comments.
	ClearanceClassified = 2

	// ClearanceConfidential users can create, edit, and delete their own posts and author debriefs.
	ClearanceConfidential = 3

	// ClearanceSecret users can author GOIs and missions, and make necessary edits and deletes to their own.
	ClearanceSecret = 4

	// ClearanceTopSecret users are admins and can do basically whatever they want.
	ClearanceTopSecret = 5
)

const (
	ActionCreate = iota
	ActionEdit   = iota
	ActionDelete = iota
)

const (
	ContentPost = iota
	ContentComment
	ContentGOI
	ContentMission
	ContentDebrief
)

func AtLeast(have, need Clearance) bool {
	return have >= need
}

func Can(
	c Clearance,
	a Action,
	ct Content,
	isOwner bool,
) bool {
	switch ct {

	case ContentComment:
		switch a {
		case ActionCreate:
			return c >= ClearanceClassified

		case ActionEdit, ActionDelete:
			if isOwner {
				return c >= ClearanceClassified
			}
			return c >= ClearanceTopSecret
		}

	case ContentPost:
		switch a {
		case ActionCreate:
			return c >= ClearanceConfidential

		case ActionEdit, ActionDelete:
			if isOwner {
				return c >= ClearanceConfidential
			}
			return c >= ClearanceTopSecret
		}

	case ContentDebrief:
		switch a {
		case ActionCreate:
			return c >= ClearanceConfidential

		case ActionEdit, ActionDelete:
			if isOwner {
				return c >= ClearanceConfidential
			}
			return c >= ClearanceTopSecret
		}

	case ContentGOI:
		switch a {
		case ActionCreate:
			return c >= ClearanceSecret

		case ActionEdit, ActionDelete:
			if isOwner {
				return c >= ClearanceSecret
			}
			return c >= ClearanceTopSecret
		}

	case ContentMission:
		switch a {
		case ActionCreate:
			return c >= ClearanceSecret

		case ActionEdit, ActionDelete:
			if isOwner {
				return c >= ClearanceSecret
			}
			return c >= ClearanceTopSecret
		}
	}

	return false
}
