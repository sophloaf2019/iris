export class Debrief {
	id: number;
	createdAt: Date;
	updatedAt: Date;
	deletedAt: Date | null;
	missionID: number
	title: string;
	slug: string;
	summary: string;
	authorID: number;
	userIDs: number[];

	constructor() {
		this.id = 0;
		this.createdAt = new Date();
		this.updatedAt = new Date();
		this.deletedAt = null;
		this.missionID = 0;
		this.title = "";
		this.slug = "";
		this.summary = "";
		this.authorID = 0;
		this.userIDs = [];
	}

	fromJSON(raw: any): this {
		this.id = raw['id'];
		this.createdAt = new Date(raw['createdAt']);
		this.updatedAt = new Date(raw['updatedAt']);
		this.deletedAt = new Date(raw['deletedAt']);
		this.missionID = raw['missionID'];
		this.title = raw['title'];
		this.slug = raw['slug'];
		this.summary = raw['summary'];
		this.authorID = raw['authorID'];
		this.userIDs = raw['userIDs'] || [];
		return this;
	}
}
