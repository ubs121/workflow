function leave(req) {
 // замууд
 if (req.Action == 'nextTriggers') {
	switch (req.FromNode) {
	case 'хүсэлт':
	  return [
	  	{name:'approve', label:'Зөвшөөрөх'}, 
	  	{name:'reject', label:'Татгалзах'}
	  	];
	default:
	  return [];
	}
 }
 // Дараагийн төлөвт шилжүүлэх
 else {
   var o = db.Leave.findOne(req.ObjectId);
   
   if (!o.Duration) {
     throw 'Чөлөөний хугацаа тодорхойгүй байна!';
   }
   
   switch (req.Action) {
   case 'approve':
     o.State = 'зөвшөөрсөн';
     // TODO: 3 хоногоос илүү бол дараагийн түвшний удирдлагад шилжих
     // return 'Чөлөөний хүсэлт ' + to_str + ' руу шилжлээ';
     break;
   case 'reject':
     o.State = 'татгалзсан';
     break;
   }
   
   try {
     db.Leave.save(o);
     return 'Дууслаа! ' + o.State;
   } catch (e) {
     throw 'Чөлөөний хүсэлт хадгалахад алдаа гарлаа: ' + e;
   }
 }
}