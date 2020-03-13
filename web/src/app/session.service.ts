// import { Injectable } from '@angular/core';
//
// @Injectable({
//   providedIn: 'root'
// })
// export class SessionService {
//   private afAuth: any;
//   private session: any;
//   private sessionSubject: any;
//   checkLogin(): void { // 追加
//     this.afAuth
//       .authState
//       .subscribe(auth => {
//         // ログイン状態を返り値の有無で判断
//         this.session.login = (!!auth);
//         this.sessionSubject.next(this.session);
//       });
//   }
// }
