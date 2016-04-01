(define-module (list-ops)
  #:export (my-length
            my-reverse
            my-map
            my-filter
            my-fold
            my-append
            my-concatenate
            ))

; 1 2 3
; -
; ((0 - 1) - 2) - 3 = -6; acc x, left -> foldl
; 3 - (2 - (1 - 0)) =  2; x acc, left -> fold
; ((0 - 3) - 2) - 1 = -6; acc x, right
; 1 - (2 - (3 - 0)) =  2; x acc, right -> foldr


; tail-rec
(define (my-fold f z l)
  (define (iter rest acc)
    (if (null? rest) acc
      (iter (cdr rest) (f (car rest) acc))
    )
  )
  (iter l z)
)

(define (my-foldl f z l)
  (my-fold (lambda (x acc) (f acc x)) z l)
)

(define (my-foldr f z l)
  (if (null? l) z (f (car l) (my-foldr f z (cdr l))))
)

; try using my-fold
(define (my-length l)
  (my-fold (lambda (x acc) (1+ acc)) 0 l)
)

(define (my-reverse l)
  (my-fold cons '() l)
)

(define (my-map f l)
  (my-foldr (lambda (x acc) (cons (f x) acc)) '() l)
)

(define (my-filter f l)
  (my-foldr (lambda (x acc) (if (f x) (cons x acc) acc)) '() l)
)

(define (my-append x y)
  (my-foldr cons y x)
)

(define (my-concatenate l)
  (my-foldl my-append '() l)
)
