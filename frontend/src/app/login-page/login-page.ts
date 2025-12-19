import { Component, inject, signal } from '@angular/core';
import { Field, form, required } from '@angular/forms/signals';
import { AuthService } from '../../services/auth';
import { Router } from '@angular/router';

@Component({
	selector: 'app-login-page',
	imports: [Field],
	templateUrl: './login-page.html',
	styleUrl: './login-page.css',
})
export class LoginPage {
	readonly loginModel = signal<{ username: string; password: string }>({
		username: '',
		password: '',
	});
	readonly loginForm = form(this.loginModel, (sch) => {
		required(sch.username, { message: 'Please provide your username.' });
		required(sch.password, { message: 'Please provide your password.' });
	});
	private authSvc: AuthService = inject(AuthService);
	private router: Router = inject(Router);

	async submitLoginForm(e: Event) {
		e.preventDefault();
		if (!this.loginForm().valid()) {
			this.loginForm().markAsTouched();
			return;
		}
		let username = this.loginForm.username().value();
		let password = this.loginForm.password().value();
		await this.authSvc.login(username, password);
		await this.router.navigate(['/app']);
	}
}
