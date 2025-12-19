import { Color } from "./../color";

export class GOI {
	id: number = 0;
	createdAt: Date = new Date();
	updatedAt: Date = new Date();
	deletedAt: Date | null = null;
	authorID: number = 0;
	name: string = "";
	slug: string = "";
	primaryColor: Color = new Color();
	secondaryColor: Color = new Color();
	mo: string = "";
	locations: Location[] = [];
	active: boolean = false;
	assignedTo: number = 0;

	fromJSON(raw: any): this {
		this.id = raw.id;
		this.createdAt = new Date(raw.createdAt);
		this.updatedAt = new Date(raw.updated_at);
		this.deletedAt = raw.deletedAt ? new Date(raw.deletedAt) : null;
		this.authorID = raw.authorID;
		this.name = raw.name;
		this.slug = raw.slug;
		this.primaryColor = raw.primaryColor;
		this.secondaryColor = raw.secondaryColor;
		this.mo = raw.mo;
		this.locations = raw.locations;
		this.active = raw.active;
		this.assignedTo = raw.assignedTo;

		return this;
	}
}
