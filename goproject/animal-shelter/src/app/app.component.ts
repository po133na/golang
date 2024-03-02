import { Component } from '@angular/core';
import { Animal } from '../animals';
import { animals } from '../animals';

@Component({
  selector: 'app-animal-list',
  template: `
    <div *ngFor="let animal of animals">
      <h2>{{ animal.name }}</h2>
      <p>{{ animal.description }}</p>
      <img [src]="animal.imageUrl" alt="{{ animal.name }}" width="200">
    </div>
  `
})
export class AnimalListComponent {
  animals: Animal[] = animals;
}
