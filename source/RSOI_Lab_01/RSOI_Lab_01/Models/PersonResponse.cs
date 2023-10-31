using RSOI_Lab_01.Entities;
using System.ComponentModel.DataAnnotations;

namespace RSOI_Lab_01.Models
{
    public class PersonResponse
    {
        /// <summary>
        /// Id юзера
        /// </summary>
        [Required]
        public int Id { get; set; }

        /// <summary>
        /// Имя юзера
        /// </summary>
        [Required]
        public string Name { get; set; }

        /// <summary>
        /// Возраст
        /// </summary>
        public int? Age { get; set; }

        /// <summary>
        /// Адрес юзера
        /// </summary>
        public string? Adress { get; set; }

        /// <summary>
        /// Работа юзера
        /// </summary>
        public string? Work { get; set; }

        public PersonResponse(int id, string name, int? age = null, string? adress = null, string? work = null)
        {
            Id = id;
            Name = name;
            Age = age;
            Adress = adress;
            Work = work;
        }

        public PersonResponse(Person person)
        {
            if (person.Id is null || person.Name is null)
                throw new ArgumentNullException(nameof(person), "Person has no id or name");

            Id = person!.Id.Value;
            Name = person!.Name;
            Age = person.Age;
            Adress = person.Adress;
            Work = person.Work;
        }
    }
}
