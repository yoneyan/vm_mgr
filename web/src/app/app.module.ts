import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {FormsModule} from "@angular/forms";

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {LoginComponent} from './login/login.component';
import {RouterModule} from "@angular/router";
import {ReactiveFormsModule} from "@angular/forms";
import {HttpClientModule} from "@angular/common/http";
import {NotfoundComponent} from './error/notfound/notfound.component';
import {AuthGuard} from "./guard/auth.guard";
import {DashboardComponent} from './dashboard/dashboard.component';
import {LogoutComponent} from './error/logout/logout.component';

import {AuthService} from './service/auth/auth.service';
import { TopBarComponent } from './bar/top-bar/top-bar.component';
import { VmComponent } from './info/vm/vm.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    NotfoundComponent,
    DashboardComponent,
    LogoutComponent,
    TopBarComponent,
    VmComponent,
    // AuthComponent,

  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    RouterModule.forRoot([
      {path: 'dashboard', component: DashboardComponent, canActivate: [AuthGuard]},
      {path: 'login', component: LoginComponent,},
      {path: 'logout', component: LogoutComponent,},
      {path: '**', component: NotfoundComponent},
    ]),
    ReactiveFormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {
}
