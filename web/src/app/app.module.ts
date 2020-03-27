import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {FormsModule, ReactiveFormsModule} from "@angular/forms";

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {LoginComponent} from './login/login.component';
import {RouterModule} from "@angular/router";
import {HttpClientModule} from "@angular/common/http";
import {NotfoundComponent} from './error/notfound/notfound.component';
import {AuthGuard} from "./guard/auth.guard";
import {DashboardComponent} from './dashboard/dashboard.component';
import {LogoutComponent} from './error/logout/logout.component';
import {TopBarComponent} from './bar/top-bar/top-bar.component';
import {VmComponent} from './info/vm/vm.component';
import {DetailVMComponent} from './dashboard/detail-vm/detail-vm.component';
import {CreateVmComponent} from './dashboard/create-vm/create-vm.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    NotfoundComponent,
    DashboardComponent,
    LogoutComponent,
    TopBarComponent,
    VmComponent,
    DetailVMComponent,
    CreateVmComponent,
    // AuthComponent,

  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    RouterModule.forRoot([
      {path: '', redirectTo: '/dashboard', pathMatch: 'full'},
      {path: 'dashboard', component: DashboardComponent, canActivate: [AuthGuard]},
      {path: 'dashboard/create', component: CreateVmComponent, canActivate: [AuthGuard]},
      {path: 'dashboard/vm/:id', component: DetailVMComponent, canActivate: [AuthGuard]},
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


export class VMModel{
  constructor(
    public nodeid: string,
    public name: string,
    public group: string,
    public cpu: number,
    public memory: number,
    public storagetype: string,
    public storage: number,
    public image: string
  ){}
}

