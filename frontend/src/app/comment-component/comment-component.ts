import { Component, EventEmitter, inject, Input, Output, signal } from '@angular/core';
import { Comment } from '../../types/social/comment';
import { AuthService } from '../../services/auth';
import { CommentService } from '../../services/social';
import { Action, Content } from '../../types/auth/user';
import { CommentFormComponent } from '../comment-form-component/comment-form-component';

@Component({
	selector: 'comment-component',
	imports: [CommentFormComponent],
	templateUrl: './comment-component.html',
	styleUrl: './comment-component.css',
})
export class CommentComponent {
	@Input() comment!: Comment;

	public commentSignal = signal<Comment>(new Comment());

	private authService: AuthService = inject(AuthService);
	private commentService: CommentService = inject(CommentService);

	@Output() commentDeleted = new EventEmitter<null>();
	@Output() commentUpdated = new EventEmitter<Comment>();

	public canEdit = signal<boolean>(false);
	public canDelete = signal<boolean>(false);
	public isEditing = signal<boolean>(false);

	public username = signal<string>('');
	public comments = signal<Comment[]>([]);

	async ngOnInit() {
		this.commentSignal.set(this.comment);
		this.username.set((await this.authService.getUserByID(this.comment.userID)).username);
		this.comments.set(await this.commentService.getForParent('comment', this.comment.id));
		let user = this.authService.user;
		if (user) {
			this.canEdit.set(
				await this.authService.userCan(
					user.id,
					Action.Edit,
					Content.Comment,
					user.id == this.comment.userID,
				),
			);
			this.canDelete.set(
				await this.authService.userCan(
					user.id,
					Action.Delete,
					Content.Comment,
					user.id == this.comment.userID,
				),
			);
		}
	}

	async editFormSubmitted(c: Comment) {
		await this.commentService.update(this.comment.id, c);
		this.commentUpdated.emit(c);
		this.isEditing.set(false);
		this.commentSignal.set(await this.commentService.getByID(this.comment.id));
	}

	async deleteButtonPressed() {
		await this.commentService.delete(this.comment.id);
		this.commentDeleted.emit();
	}
}
