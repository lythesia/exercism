(define-module (anagram)
  #:export (anagrams-for))

(define (anagrams-for wrd lst)
  (define (conv w)
    (sort (map char-downcase (string->list w)) char<?)
  )
  (let ((ref (conv wrd)))
    (filter
      (lambda (e) (and (equal? ref (conv e)) (not (equal? wrd e))))
      lst
    )
  )
)
