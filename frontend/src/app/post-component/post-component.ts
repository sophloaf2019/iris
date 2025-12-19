import { Component, EventEmitter, inject, Input, Output, signal } from '@angular/core';
import { Post } from '../../types/social/post';
import { Comment } from '../../types/social/comment';
import { CommentService, PostService } from '../../services/social';
import { AuthService } from '../../services/auth';
import { DatePipe } from '@angular/common';
import { CommentComponent } from '../comment-component/comment-component';
import { PostFormComponent } from '../post-form-component/post-form-component';
import { Action, Content } from '../../types/auth/user';
import { CommentFormComponent } from '../comment-form-component/comment-form-component';

@Component({
	selector: 'post-component',
	imports: [DatePipe, CommentComponent, PostFormComponent, CommentFormComponent],
	templateUrl: './post-component.html',
	styleUrl: './post-component.css',
})
export class PostComponent {
	// The Post this component will display
	@Input() post!: Post;
	// The events emitted on updates or deletes respectively
	@Output() postUpdated = new EventEmitter<Post>();
	@Output() postDeleted = new EventEmitter<null>();

	// Relevant services for enriching data (user, comment)
	private authService: AuthService = inject(AuthService);
	private commentService: CommentService = inject(CommentService);
	private postService: PostService = inject(PostService);

	public username = signal<string>('');
	public comments = signal<Comment[]>([]);

	public canEdit = signal<boolean>(false);
	public canDelete = signal<boolean>(false);
	public isEditing = signal<boolean>(false);
	public canComment = signal<boolean>(false);

	async ngOnInit() {
		this.refreshComments();
		this.username.set((await this.authService.getUserByID(this.post.userID)).username);
		let user = this.authService.user;
		if (user) {
			this.canEdit.set(
				await this.authService.userCan(
					user.id,
					Action.Edit,
					Content.Post,
					user.id == this.post.id,
				),
			);
			this.canDelete.set(
				await this.authService.userCan(
					user.id,
					Action.Delete,
					Content.Post,
					user.id == this.post.id,
				),
			);
			this.canComment.set(
				await this.authService.userCan(
					user.id,
					Action.Create,
					Content.Comment,
					user.id == this.post.id,
				),
			);
		}
	}

	async refreshComments() {
		this.comments.set([]);
		this.comments.set(await this.commentService.getForParent('post', this.post.id));
	}

	async editFormSubmitted(p: Post) {
		await this.postService.update(this.post.id, p);
		this.postUpdated.emit(p);
	}

	async deleteButtonClicked() {
		await this.postService.delete(this.post.id);
		this.postDeleted.emit();
	}

	async postComment(c: Comment) {
		await this.commentService.create(c);
		this.postUpdated.emit();
	}
}
