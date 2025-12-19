import { Component, EventEmitter, Input, Output, signal } from '@angular/core';
import { Comment } from '../../types/social/comment';
import { form, required, Field } from '@angular/forms/signals';

@Component({
	selector: 'comment-form-component',
	imports: [Field],
	templateUrl: './comment-form-component.html',
	styleUrl: './comment-form-component.css',
})
export class CommentFormComponent {
	@Input() parent_type!: string;
	@Input() parent_id!: number;
	@Input() comment?: Comment;

	public cSig = signal<Comment>(new Comment());
	public cForm = form(this.cSig, (sch) => {
		required(sch.message);
	});

	ngOnInit() {
		if (this.comment) {
			this.cSig.set(this.comment);
		}
	}

	@Output() emitChanges = new EventEmitter<Comment>();

	formSubmitted(e: Event) {
		e.preventDefault();
		if (this.cForm().valid() == false) {
			return;
		}
		let comment = this.cForm().value();
		comment.parentID = this.parent_id;
		comment.parentType = this.parent_type;
		this.emitChanges.emit(comment);
	}

	wipe() {
		this.cSig.set(new Comment());
		this.cForm().reset();
	}
}
