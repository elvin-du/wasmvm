(module
 (table 0 anyfunc)
 (memory $0 1)
 (export "memory" (memory $0))
 (export "_Z3addii" (func $_Z3addii))
 (export "_Z3sumff" (func $_Z3sumff))
 (func $_Z3addii (; 0 ;) (param $0 i32) (param $1 i32) (result i32)
  (i32.add
   (get_local $1)
   (get_local $0)
  )
 )
 (func $_Z3sumff (; 1 ;) (param $0 f32) (param $1 f32) (result f32)
  (f32.add
   (get_local $0)
   (get_local $1)
  )
 )
)
