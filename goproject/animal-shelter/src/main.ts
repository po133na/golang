import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/app.config';
import {AnimalListComponent } from './app/app.component';

bootstrapApplication(AnimalListComponent, appConfig)
  .catch((err) => console.error(err));
