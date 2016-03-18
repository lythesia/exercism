(define-module (robot)
  #:export (build-robot
            robot-name
            reset-name)
  #:autoload (srfi srfi-1) (iota))

(define (build-robot)
  (define (gen-name)
    (define (gen-upper)
      (integer->char (+ 65 (random 26)))
    )
    (string-append
      (string (gen-upper) (gen-upper))
      (string-take-right (symbol->string (gensym)) 3)
    )
  )
  (let ((name (gen-name)))
    (define (dispatch m)
      (cond
        ((eq? m 'get-name) name)
        ((eq? m 'reset-name) (set! name (gen-name)))
      )
    )
    dispatch
  )
)

(define (robot-name r) (r 'get-name))
(define (reset-name r) (r 'reset-name))
