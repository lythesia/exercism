(define-module (nucleotide-count)
  #:export (nucleotide-counts dna-count))

(define (dna-count t s)
  (if (memq t '(#\A #\C #\G #\T))
    (string-count s t)
    (error "invalid nucleotide" t)
  )
)

(use-modules (ice-9 hash-table))

(define (nucleotide-counts s)
  (let* ((result (make-hash-table 4)))
    (for-each (lambda (c) (hash-set! result c 0)) '(#\A #\C #\G #\T))
    (string-for-each
      (lambda (c) (hash-set! result c (1+ (hash-ref result c))))
      s
    )
    (sort (hash-map->list cons result) (lambda (x y) (char<? (car x) (car y))))
  )
)
