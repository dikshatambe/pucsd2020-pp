import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { SummaryComponent } from './summary/summary.component';
import { TechnicalSkillComponent } from './technical-skill/technical-skill.component';
import { EducationComponent } from './education/education.component';
import { ExperienceComponent } from './experience/experience.component';
import { LanguageComponent } from './language/language.component';


const routes: Routes = [
  { path: 'SummaryComponent', component: SummaryComponent },
  { path: 'TechnicalSkillComponent', component: TechnicalSkillComponent },
  { path: 'EducationComponent', component: EducationComponent },
  { path: 'ExperienceComponent', component: ExperienceComponent },
  { path: 'LanguageComponent', component: LanguageComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
