import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { EventDashboardComponent } from './event-dashboard/event-dashboard.component';
import { EventScheduleComponent } from './event-schedule/event-schedule.component';
import { EventManageComponent } from './event-manage/event-manage.component';

@NgModule({
  declarations: [
    AppComponent,
    EventDashboardComponent,
    EventScheduleComponent,
    EventManageComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
