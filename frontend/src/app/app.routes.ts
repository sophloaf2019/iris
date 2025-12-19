import { Routes } from '@angular/router';
import { LoginPage } from './login-page/login-page';
import { AppPage } from './app-page/app-page';
import { AnnouncementPage } from './announcement-page/announcement-page';

export const routes: Routes = [
	{
		path: '',
		component: LoginPage,
	},
	{
		path: 'app',
		component: AppPage,
		children: [{ path: '', component: AnnouncementPage }],
	},
];
