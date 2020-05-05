import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AddComponent } from './add/add.component';
import { SearchComponent } from './search/search.component';
import { UpdateComponent } from './update/update.component';
import { DeleteComponent } from './delete/delete.component';


const routes: Routes = [
  { path: 'AddComponent', component: AddComponent },
  { path: 'SearchComponent', component: SearchComponent },
  { path: 'UpdateComponent', component: UpdateComponent },
  { path: 'DeleteComponent', component: DeleteComponent }

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
