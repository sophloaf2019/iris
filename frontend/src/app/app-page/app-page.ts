import { Component } from '@angular/core';
import { RouterOutlet, RouterLinkWithHref, RouterLink } from '@angular/router';

@Component({
	selector: 'app-page',
	imports: [RouterOutlet, RouterLink],
	templateUrl: './app-page.html',
	styleUrl: './app-page.css',
})
export class AppPage {}
