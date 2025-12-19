export class Post {
	id: number = 0;
	createdAt: Date = new Date(0);
	updatedAt: Date = new Date(0);
	deletedAt: Date | null = null;
	title: string = '';
	userID: number = 0;
	message: string = '';

	fromJSON(raw: any): this {
		this.id = raw.id;
		this.createdAt = new Date(raw.createdAt);
		this.updatedAt = new Date(raw.updatedAt);
		this.deletedAt = raw.deletedAt ? new Date(raw.deletedAt) : null;
		this.title = raw.title;
		this.userID = raw.userID;
		this.message = raw.message;

		return this;
	}
}
