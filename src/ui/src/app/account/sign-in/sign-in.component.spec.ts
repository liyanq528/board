/**
 @fileName: sign-in.component.spec
 @author: liyanq
 @dateTime: 25/01/2018 13:16
 @desc:
 */
import { By } from '@angular/platform-browser';
import { DebugElement, Injectable } from "@angular/core";
import { ActivatedRoute, NavigationExtras, Router } from "@angular/router";
import { async, ComponentFixture, fakeAsync, TestBed, tick } from '@angular/core/testing';
import { SignInComponent } from "./sign-in.component"
import { AccountService } from "../account.service";
import { AppModule } from "../../app.module";
import { AppComponent } from "../../app.component";
import { AppInitService } from "app/app.init.service";
import Spy = jasmine.Spy;

export function newEvent(eventName: string, bubbles = false, cancelable = false) {
  let evt = document.createEvent('CustomEvent');  // MUST be 'CustomEvent'
  evt.initCustomEvent(eventName, bubbles, cancelable, null);
  return evt;
}

@Injectable()
export class RouterStub {
  navigate(commands: any[], extras?: NavigationExtras) {
  }
}

describe("SignInComponent", () => {
  const fakeActivatedRoute = {
    snapshot: {
      data: {
        systeminfo: {
          auth_mode: "ldap_auth",
          board_host: "10.165.22.61",
          board_version: "dev",
          init_project_repo: "created",
          redirection_url: "http://redirection.mydomain.com",
          set_auth_password: "updated",
          sync_k8s: "created"
        }
      }
    }
  };
  let fixture: ComponentFixture<SignInComponent>;
  let component: SignInComponent;
  let accountService: AccountService;
  let signInSpy: Spy;
  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [AppModule],
      providers: [
        {provide: Router, useClass: RouterStub},
        {provide: ActivatedRoute, useValue: fakeActivatedRoute}]
    }).compileComponents().then(() => {
      /*The AppComponent must be first create,
       because set currentLang on it's constructor function*/
      TestBed.createComponent(AppComponent);
      fixture = TestBed.createComponent(SignInComponent);
      fixture.detectChanges();//trigger initial data binding
      component = fixture.componentInstance;
      accountService = fixture.debugElement.injector.get(AccountService);
      signInSpy = spyOn(accountService, "signIn");
    })
  }));

  function testErrorHandle(status: number): void {
    let btnSubmitDebug: DebugElement = fixture.debugElement.query(By.css("[type='submit']"));
    let btnSubmit: HTMLInputElement = btnSubmitDebug.nativeElement as HTMLInputElement;
    let appInitService = fixture.debugElement.injector.get(AppInitService);
    component.signInUser.username = "hello";
    component.signInUser.password = "world";
    signInSpy.and.returnValue(Promise.reject({status: status}));
    btnSubmit.click();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      setTimeout(() => {
        expect(signInSpy.calls.any()).toBe(true, '/v1/api/sign-in restApi called');
        expect(appInitService.token).toBeUndefined();
      }, 500);
    });
  }

  it('should create the SignInComponent', () => {
    expect(SignInComponent).toBeTruthy();
  });

  it('should submit button is disabled', () => {
    let btnSubmitDebug: DebugElement = fixture.debugElement.query(By.css("[type='submit']"));
    let btnSubmit: HTMLInputElement = btnSubmitDebug.nativeElement as HTMLInputElement;
    fixture.detectChanges();
    expect<boolean>(btnSubmit.disabled).toBeTruthy("submit button is disabled");
  });

  it('should submit button is enable', fakeAsync(() => {
    let btnSubmitDebug: DebugElement = fixture.debugElement.query(By.css("[type='submit']"));
    let btnSubmit: HTMLInputElement = btnSubmitDebug.nativeElement as HTMLInputElement;
    component.signInUser.username = "hello";
    component.signInUser.password = "world";
    tick();
    expect<boolean>(btnSubmit.disabled).toBeFalsy("submit button is enabled");
  }));

  it("test sign-in click and catch 400 error code", async(() => {
    testErrorHandle(400);
  }));

  it("test sign-in click and catch 409 error code", async(() => {
    testErrorHandle(409);
  }));

  it("should sign-in with {username:'admin',password:'xxxxxx' success}", async(() => {
    let btnSubmitDebug: DebugElement = fixture.debugElement.query(By.css("[type='submit']"));
    let btnSubmit: HTMLInputElement = btnSubmitDebug.nativeElement as HTMLInputElement;
    let appInitService = fixture.debugElement.injector.get(AppInitService);
    let router = fixture.debugElement.injector.get(Router);
    let navSpy = spyOn(router, 'navigate');
    component.signInUser.username = "admin";
    component.signInUser.password = "xxxxxx";
    signInSpy.and.returnValue(Promise.resolve({token: 'tokenString'}));
    btnSubmit.click();
    fixture.whenStable().then(() => {
      fixture.detectChanges();
      const navArgs = navSpy.calls.first().args[0];
      expect(signInSpy.calls.any()).toBe(true, '/v1/api/sign-in restApi called');
      expect(navSpy.calls.any()).toBe(true, 'navigate called');
      expect(navArgs[0]).toContain('dashboard', 'nav to dashboard detail URL');
      expect(appInitService.token).toBe("tokenString", 'appInitService.token is undefined')
    });
  }))
});