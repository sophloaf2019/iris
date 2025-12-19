import { Component, inject, signal, ViewChild } from '@angular/core';
import { PostService } from '../../services/social';
import { Post } from '../../types/social/post';
import { PostComponent } from '../post-component/post-component';
import { PostFormComponent } from '../post-form-component/post-form-component';
import { AuthService } from '../../services/auth';
import { Action, Content } from '../../types/auth/user';

@Component({
	selector: 'app-announcement-page',
	imports: [PostComponent, PostFormComponent],
	templateUrl: './announcement-page.html',
	styleUrl: './announcement-page.css',
})
export class AnnouncementPage {
	private postService: PostService = inject(PostService);
	private authService: AuthService = inject(AuthService);
	public posts = signal<Post[]>([]);
	public canPost = signal<boolean>(false);

	@ViewChild(PostFormComponent) newForm!: PostFormComponent;

	async ngOnInit() {
		this.posts.set(await this.postService.list());
		let user = this.authService.user;
		if (user) {
			this.canPost.set(
				await this.authService.userCan(user?.id, Action.Create, Content.Post, false),
			);
		}
	}

	async submitNewPost(p: Post) {
		// TODO
		// this is weird
		// fix it later
		// - sophie
		await this.postService.create(p);
		this.refresh();
		this.newForm.wipe();
	}

	async refresh() {
		this.posts.set([]);
		this.posts.set(await this.postService.list());
	}
}
