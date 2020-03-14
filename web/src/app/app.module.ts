import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {FormsModule} from "@angular/forms";

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {LoginComponent} from './login/login.component';
import {RouterModule} from "@angular/router";
import {ReactiveFormsModule} from "@angular/forms";
import {Test} from './test';
import {HttpClientModule} from "@angular/common/http";
import {NotfoundComponent} from './error/notfound/notfound.component';
import {AuthGuard} from "./guard/auth.guard";
import { DashboardComponent } from './dashboard/dashboard.component';

// import { AuthComponent } from './service/auth/auth.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    NotfoundComponent,
    DashboardComponent,
    // AuthComponent,

  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    RouterModule.forRoot([
      {path: '', component: DashboardComponent, canActivate: [AuthGuard]},
      {path: 'login', component: LoginComponent,},
      {path: '**', component: NotfoundComponent},
    ]),
    ReactiveFormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {
}
