export class Mission {
	id: number = 0;
	createdAt: Date = new Date();
 	updatedAt: Date = new Date();
	deletedAt: Date | null = null;
	title: string = "";
	slug: string = "";
	briefing: string = "";
	tags: string[] = [];
	goiID: number = 0;
	authorID: number = 0;
	location: Location = new Location();
	interestedUsers: number[] = [];

	fromJSON(raw: any): this {
		this.id = raw.id;
		this.createdAt = new Date(raw.createdAt);
		this.updatedAt = new Date(raw.updatedAt);
		this.deletedAt = new Date(raw.deletedAt);
		this.title = raw.title;
		this.slug = raw.slug;
		this.briefing = raw.briefing;
		this.tags = raw.tags;
		this.goiID = raw.goiID;
		this.authorID = raw.authorID;
		this.location = raw.location;
		this.interestedUsers = raw.interestedUsers;
		return this;
	}
}
