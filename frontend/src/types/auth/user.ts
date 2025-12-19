export class User {
	id: number;
	createdAt: Date;
	updatedAt: Date;
	deletedAt: Date | null;
	username: string;
	clearance: number;

	constructor() {
		this.id = 0;
		this.createdAt = new Date();
		this.updatedAt = new Date();
		this.deletedAt = null;
		this.username = '';
		this.clearance = 0;
	}

	fromJSON(raw: any): this {
		this.id = raw['id'];
		this.createdAt = new Date(raw['createdAt']);
		this.updatedAt = new Date(raw['updatedAt']);
		this.deletedAt = new Date(raw['deletedAt']);
		this.username = raw['username'];
		this.clearance = raw['clearance'];
		return this;
	}
}

export enum Action {
	Create = 0,
	Edit,
	Delete,
}

export enum Content {
	Post = 0,
	Comment,
	GOI,
	Mission,
	Debrief,
}
