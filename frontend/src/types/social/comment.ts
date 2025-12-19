export class Comment {
	id: number = 0;
	createdAt: Date = new Date(0);
	updatedAt: Date = new Date(0);
	deletedAt: Date | null = null;
	userID: number = 0;
	message: string = '';
	parentID: number = 0;
	parentType: string = '';

	static fromJSON(raw: any): Comment {
		const c = new Comment();

		c.id = raw.id;
		c.createdAt = new Date(raw.createdAt);
		c.updatedAt = new Date(raw.updatedAt);
		c.deletedAt = raw.deletedAt ? new Date(raw.deletedAt) : null;
		c.userID = raw.userID;
		c.message = raw.message;
		c.parentID = raw.parentID;
		c.parentType = raw.parentType;

		return c;
	}
}
