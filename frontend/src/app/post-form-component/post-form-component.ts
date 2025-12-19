import { Component, EventEmitter, Input, Output, signal } from '@angular/core';
import { Post } from '../../types/social/post';
import { form, Field, required } from '@angular/forms/signals';

@Component({
	selector: 'post-form-component',
	imports: [Field],
	templateUrl: './post-form-component.html',
	styleUrl: './post-form-component.css',
})
export class PostFormComponent {
	@Input() post?: Post;
	public postSignal = signal<Post>(new Post());
	public postForm = form(this.postSignal, (sch) => {
		required(sch.title, { message: 'Please provide a title for your post.' });
		required(sch.message, { message: 'Please provide a body for your post.' });
	});

	@Output() emitChanges = new EventEmitter<Post>();

	ngOnInit() {
		if (this.post != null) {
			this.postSignal.set(this.post);
		}
		this.postForm().reset();
	}

	submit(e: Event) {
		e.preventDefault();
		if (this.postForm().valid() == false) {
			this.postForm().markAsTouched();
			return;
		}
		this.emitChanges.emit(this.postForm().value());
	}

	wipe() {
		this.postSignal.set(new Post());
		this.postForm().reset();
	}
}
